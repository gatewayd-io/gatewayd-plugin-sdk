package postgres

import (
	"github.com/jackc/pgx/v5/pgproto3"
)

// BuildAuthCleartextChallenge builds an AuthenticationCleartextPassword message.
func BuildAuthCleartextChallenge() ([]byte, error) {
	msg := &pgproto3.AuthenticationCleartextPassword{}
	return msg.Encode(nil)
}

// BuildAuthMD5Challenge builds an AuthenticationMD5Password message with the given salt.
func BuildAuthMD5Challenge(salt [4]byte) ([]byte, error) {
	msg := &pgproto3.AuthenticationMD5Password{Salt: salt}
	return msg.Encode(nil)
}

// BuildAuthSASLChallenge builds an AuthenticationSASL message listing the given mechanisms.
func BuildAuthSASLChallenge(mechanisms []string) ([]byte, error) {
	msg := &pgproto3.AuthenticationSASL{AuthMechanisms: mechanisms}
	return msg.Encode(nil)
}

// BuildAuthSASLContinue builds an AuthenticationSASLContinue message with server-first data.
func BuildAuthSASLContinue(data []byte) ([]byte, error) {
	msg := &pgproto3.AuthenticationSASLContinue{Data: data}
	return msg.Encode(nil)
}

// BuildAuthSASLFinal builds an AuthenticationSASLFinal message with server-final data.
func BuildAuthSASLFinal(data []byte) ([]byte, error) {
	msg := &pgproto3.AuthenticationSASLFinal{Data: data}
	return msg.Encode(nil)
}

// BuildAuthOk builds the full post-authentication response sequence:
// AuthenticationOk + ParameterStatus(server_version) + ParameterStatus(client_encoding)
// + BackendKeyData + ReadyForQuery.
func BuildAuthOk(serverVersion string, processID uint32, secretKey uint32) ([]byte, error) {
	authOK := &pgproto3.AuthenticationOk{}
	buf, err := authOK.Encode(nil)
	if err != nil {
		return nil, err
	}

	pStatVersion := &pgproto3.ParameterStatus{
		Name:  "server_version",
		Value: serverVersion,
	}
	buf, err = pStatVersion.Encode(buf)
	if err != nil {
		return nil, err
	}

	pStatEncoding := &pgproto3.ParameterStatus{
		Name:  "client_encoding",
		Value: "UTF8",
	}
	buf, err = pStatEncoding.Encode(buf)
	if err != nil {
		return nil, err
	}

	backendKeyData := &pgproto3.BackendKeyData{
		ProcessID: processID,
		SecretKey: secretKey,
	}
	buf, err = backendKeyData.Encode(buf)
	if err != nil {
		return nil, err
	}

	readyForQuery := &pgproto3.ReadyForQuery{TxStatus: 'I'}
	buf, err = readyForQuery.Encode(buf)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// BuildTerminateWithError builds a Terminate message prepended with an ErrorResponse.
func BuildTerminateWithError(msg, severity, code, detail string) ([]byte, error) {
	errResp := ErrorResponse(msg, severity, code, detail)
	terminate := &pgproto3.Terminate{}
	return terminate.Encode(errResp)
}
