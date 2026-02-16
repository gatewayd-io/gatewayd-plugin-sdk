package postgres

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"

	v1 "github.com/gatewayd-io/gatewayd-plugin-sdk/plugin/v1"
	"github.com/hashicorp/go-hclog"
	"github.com/jackc/pgx/v5/pgproto3"
)

var PostgreSQLFrontEndMessageTypes = []uint8{
	'B',
	'C',
	'D',
	'E',
	'F',
	'f',
	'd',
	'c',
	'H',
	'P',
	'p',
	'Q',
	'S',
	'X',
}

// IsPostgresStartupMessage returns true if the message is a startup message.
func IsPostgresStartupMessage(b []byte) bool {
	for _, v := range PostgreSQLFrontEndMessageTypes {
		if len(b) != 0 && b[0] == v {
			return false
		}
	}

	return true
}

func ErrorResponse(msg, severity, code, detail string) []byte {
	// NOTE: The error from the Encode method can be safely ignored because
	// the result will be nil on error. The parameters MUST be provided by us
	// and they MUST NOT be filled with user input.
	errResp, _ := (&pgproto3.ErrorResponse{
		Severity: severity,
		Message:  msg,
		Code:     code,
		Detail:   detail,
	}).Encode(nil)
	return errResp
}

// IsPostgresSSLRequest returns true if the message is a SSL request.
func IsPostgresSSLRequest(b []byte) bool {
	if len(b) < 8 {
		return false
	}

	if binary.BigEndian.Uint32(b[0:4]) != 8 {
		return false
	}

	if binary.BigEndian.Uint32(b[4:8]) != 80877103 {
		return false
	}

	return true
}

// DecodeBytes decodes a base64 encoded string to bytes.
func DecodeBytes(encoded string) []byte {
	valueBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return []byte{}
	}
	return valueBytes
}

// HandleClientMessage handles the client message.
// This function should be called from onTrafficFromClient hook.
func HandleClientMessage(req *v1.Struct, logger hclog.Logger) (*v1.Struct, error) {
	request := req.Fields["request"].GetBytesValue()
	if len(request) == 0 {
		// Nothing to parse (e.g. client disconnected).
		return req, nil
	}
	pgBackend := pgproto3.NewBackend(bytes.NewReader(request), nil)

	if IsPostgresStartupMessage(request) {
		// Handle startup message
		startupMessage, err := pgBackend.ReceiveStartupMessage()
		if err != nil {
			logger.Error("Cannot receive startup message", "error", err)
			return req, nil
		}

		switch startupMessage := startupMessage.(type) {
		case *pgproto3.StartupMessage:
			startupMsgJson, err := startupMessage.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal startup message", "error", err)
			}
			req.Fields["startupMessage"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(startupMsgJson))
		case *pgproto3.SSLRequest:
			sslRequestJson, err := startupMessage.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal SSLRequest", "error", err)
			}
			req.Fields["sslRequest"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(sslRequestJson))
		case *pgproto3.SASLInitialResponse:
			saslInitialResponseJson, err := startupMessage.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal SASLInitialResponse", "error", err)
			}
			req.Fields["saslInitialResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(saslInitialResponseJson))
		case *pgproto3.CancelRequest:
			cancelRequestJson, err := startupMessage.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CancelRequest", "error", err)
			}
			req.Fields["cancelRequest"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(cancelRequestJson))
		default:
			logger.Debug("Unknown startup message", "message", startupMessage)
		}
	} else {
		// Handle other messages
		message, err := pgBackend.Receive()
		if err != nil && !errors.Is(err, io.ErrUnexpectedEOF) && !errors.Is(err, io.EOF) {
			// Truly an invalid message
			logger.Error("Cannot receive message", "error", err)
			return req, nil
		}

		switch message := message.(type) {
		case *pgproto3.SASLResponse:
			saslResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal SASLResponse", "error", err)
			}
			req.Fields["saslResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(saslResponseJson))
		case *pgproto3.GSSResponse:
			gssResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal GSSResponse", "error", err)
			}
			req.Fields["gssResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(gssResponseJson))
		case *pgproto3.Query:
			queryJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Query", "error", err)
			}
			req.Fields["query"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(queryJson))
		case *pgproto3.Parse:
			parseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Parse", "error", err)
			}
			req.Fields["parse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(parseJson))
		case *pgproto3.Bind:
			bindJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Bind", "error", err)
			}
			req.Fields["bind"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(bindJson))
		case *pgproto3.Describe:
			describeJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Describe", "error", err)
			}
			req.Fields["describe"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(describeJson))
		case *pgproto3.Execute:
			executeJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Execute", "error", err)
			}
			req.Fields["execute"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(executeJson))
		case *pgproto3.Sync:
			syncJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Sync", "error", err)
			}
			req.Fields["sync"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(syncJson))
		case *pgproto3.Terminate:
			terminateJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Terminate", "error", err)
			}
			req.Fields["terminate"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(terminateJson))
		case *pgproto3.Close:
			closeJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Close", "error", err)
			}
			req.Fields["close"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(closeJson))
		case *pgproto3.CopyData:
			copyDataJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyData", "error", err)
			}
			req.Fields["copyData"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyDataJson))
		case *pgproto3.CopyDone:
			copyDoneJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyDone", "error", err)
			}
			req.Fields["copyDone"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyDoneJson))
		case *pgproto3.CopyFail:
			copyFailJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyFail", "error", err)
			}
			req.Fields["copyFail"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyFailJson))
		case *pgproto3.Flush:
			flushJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal Flush", "error", err)
			}
			req.Fields["flush"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(flushJson))
		case *pgproto3.CancelRequest:
			cancelRequestJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CancelRequest", "error", err)
			}
			req.Fields["cancelRequest"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(cancelRequestJson))
		case *pgproto3.PasswordMessage:
			passwordMessageJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal PasswordMessage", "error", err)
			}
			req.Fields["passwordMessage"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(passwordMessageJson))
		case *pgproto3.SASLInitialResponse:
			saslInitialResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal SASLInitialResponse", "error", err)
			}
			req.Fields["saslInitialResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(saslInitialResponseJson))
		default:
			logger.Debug("Unknown message type from client")
		}
	}

	return req, nil
}

// HandleServerMessage handles the server response.
// This function should be called from onTrafficFromServer hook.
func HandleServerMessage(resp *v1.Struct, logger hclog.Logger) (*v1.Struct, error) {
	response := resp.Fields["response"].GetBytesValue()
	pgFrontend := pgproto3.NewFrontend(bytes.NewReader(response), nil)

	for {
		message, err := pgFrontend.Receive()
		if err != nil {
			return resp, nil
		}

		if message == nil {
			return resp, nil
		}

		switch message := message.(type) {
		case *pgproto3.AuthenticationOk:
			authenticationOkJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationOk", "error", err)
			}
			resp.Fields["authenticationOk"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationOkJson))
		case *pgproto3.AuthenticationCleartextPassword:
			authenticationCleartextPasswordJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationCleartextPassword", "error", err)
			}
			resp.Fields["authenticationCleartextPassword"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationCleartextPasswordJson))
		case *pgproto3.AuthenticationMD5Password:
			authenticationMD5PasswordJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationMD5Password", "error", err)
			}
			resp.Fields["authenticationMD5Password"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationMD5PasswordJson))
		case *pgproto3.AuthenticationGSS:
			authenticationGSSJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationGSS", "error", err)
			}
			resp.Fields["authenticationGSS"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationGSSJson))
		case *pgproto3.AuthenticationSASL:
			authenticationSASLJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationSASL", "error", err)
			}
			resp.Fields["authenticationSASL"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationSASLJson))
		case *pgproto3.AuthenticationSASLContinue:
			authenticationSASLContinueJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationSASLContinue", "error", err)
			}
			resp.Fields["authenticationSASLContinue"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationSASLContinueJson))
		case *pgproto3.AuthenticationSASLFinal:
			authenticationSASLFinalJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal AuthenticationSASLFinal", "error", err)
			}
			resp.Fields["authenticationSASLFinal"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(authenticationSASLFinalJson))
		case *pgproto3.BackendKeyData:
			backendKeyDataJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal BackendKeyData", "error", err)
			}
			resp.Fields["backendKeyData"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(backendKeyDataJson))
		case *pgproto3.BindComplete:
			bindCompleteJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal BindComplete", "error", err)
			}
			resp.Fields["bindComplete"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(bindCompleteJson))
		case *pgproto3.CloseComplete:
			closeCompleteJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CloseComplete", "error", err)
			}
			resp.Fields["closeComplete"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(closeCompleteJson))
		case *pgproto3.CommandComplete:
			commandCompleteJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CommandComplete", "error", err)
			}
			resp.Fields["commandComplete"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(commandCompleteJson))
		case *pgproto3.CopyData:
			copyDataJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyData", "error", err)
			}
			resp.Fields["copyData"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyDataJson))
		case *pgproto3.CopyDone:
			copyDoneJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyDone", "error", err)
			}
			resp.Fields["copyDone"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyDoneJson))
		case *pgproto3.CopyInResponse:
			copyInResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyInResponse", "error", err)
			}
			resp.Fields["copyInResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyInResponseJson))
		case *pgproto3.CopyOutResponse:
			copyOutResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyOutResponse", "error", err)
			}
			resp.Fields["copyOutResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyOutResponseJson))
		case *pgproto3.CopyBothResponse:
			copyBothResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal CopyBothResponse", "error", err)
			}
			resp.Fields["copyBothResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(copyBothResponseJson))
		case *pgproto3.RowDescription:
			rowDescriptionJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal RowDescription", "error", err)
			}
			resp.Fields["rowDescription"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(rowDescriptionJson))
		case *pgproto3.DataRow:
			dataRowJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal DataRow", "error", err)
			}
			if resp.Fields["dataRow"] == nil {
				// Create a new list and add the first value
				resp.Fields["dataRow"] = v1.NewListValue(
					&v1.ListValue{
						Values: []*v1.Value{
							v1.NewBytesValue(dataRowJson),
						},
					},
				)
			} else {
				// Append to the existing list
				resp.Fields["dataRow"].GetListValue().Values = append(
					resp.Fields["dataRow"].GetListValue().Values,
					v1.NewBytesValue(dataRowJson),
				)
			}
		case *pgproto3.EmptyQueryResponse:
			emptyQueryResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal EmptyQueryResponse", "error", err)
			}
			resp.Fields["emptyQueryResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(emptyQueryResponseJson))
		case *pgproto3.ErrorResponse:
			errorResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal ErrorResponse", "error", err)
			}
			resp.Fields["errorResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(errorResponseJson))
		case *pgproto3.FunctionCallResponse:
			functionCallResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal FunctionCallResponse", "error", err)
			}
			resp.Fields["functionCallResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(functionCallResponseJson))
		case *pgproto3.NoData:
			noDataJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal NoData", "error", err)
			}
			resp.Fields["noData"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(noDataJson))
		case *pgproto3.NoticeResponse:
			noticeResponse := map[string]interface{}{
				"severity":            message.Severity,
				"severityUnlocalized": message.SeverityUnlocalized,
				"code":                message.Code,
				"message":             message.Message,
				"detail":              message.Detail,
				"hint":                message.Hint,
				"position":            message.Position,
				"internalPosition":    message.InternalPosition,
				"internalQuery":       message.InternalQuery,
				"where":               message.Where,
				"schemaName":          message.SchemaName,
				"tableName":           message.TableName,
				"columnName":          message.ColumnName,
				"dataTypeName":        message.DataTypeName,
				"constraintName":      message.ConstraintName,
				"file":                message.File,
				"line":                message.Line,
				"routine":             message.Routine,
				"unknownFields":       message.UnknownFields,
			}
			noticeResponseJson, err := json.Marshal(noticeResponse)
			if err != nil {
				logger.Error("Cannot marshal NoticeResponse", "error", err)
			}
			resp.Fields["noticeResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(noticeResponseJson))
		case *pgproto3.NotificationResponse:
			notificationResponseJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal NotificationResponse", "error", err)
			}
			resp.Fields["notificationResponse"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(notificationResponseJson))
		case *pgproto3.ParameterDescription:
			parameterDescriptionJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal ParameterDescription", "error", err)
			}
			resp.Fields["parameterDescription"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(parameterDescriptionJson))
		case *pgproto3.ParameterStatus:
			parameterStatusJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal ParameterStatus", "error", err)
			}
			resp.Fields["parameterStatus"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(parameterStatusJson))
		case *pgproto3.ParseComplete:
			parseCompleteJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal ParseComplete", "error", err)
			}
			resp.Fields["parseComplete"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(parseCompleteJson))
		case *pgproto3.PortalSuspended:
			portalSuspendedJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal PortalSuspended", "error", err)
			}
			resp.Fields["portalSuspended"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(portalSuspendedJson))
		case *pgproto3.ReadyForQuery:
			readyForQueryJson, err := message.MarshalJSON()
			if err != nil {
				logger.Error("Cannot marshal ReadyForQuery", "error", err)
			}
			resp.Fields["readyForQuery"] = v1.NewStringValue(
				base64.StdEncoding.EncodeToString(readyForQueryJson))
		default:
			logger.Error("Unknown message type from server")
			// TODO: return error
			return resp, nil
		}
	}
}
