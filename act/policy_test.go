package act

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPolicy_Eval(t *testing.T) {
	t.Run("it should eval program correctly", func(t *testing.T) {
		policy := MustNewPolicy("test", "1 + 1", map[string]any{})
		resp, err := policy.Eval(context.Background(), Input{})
		require.NoError(t, err)
		assert.Equal(t, 2, resp)
	})

	t.Run("it should stop eval if context is timed out", func(t *testing.T) {
		policy := MustNewPolicy("test", "repeat(\"Hi\", 100000)", map[string]any{})
		ctx, cancel := context.WithTimeout(context.Background(), 10)
		defer cancel()
		_, err := policy.Eval(ctx, Input{})
		require.Error(t, err)
	})

	t.Run("it should fail because of runtime error", func(t *testing.T) {
		policy := MustNewPolicy("test", `nonExistantVariable`, map[string]any{})
		_, err := policy.Eval(context.Background(), Input{
			Policy: map[string]any{"key": "value"},
		})
		require.ErrorIs(t, err, context.DeadlineExceeded)
	})
}

func TestPolicy_MustCompile(t *testing.T) {
	t.Run("it should compile policy correctly", func(t *testing.T) {
		policy := Policy{
			Name:     "test",
			Policy:   "1 + 1",
			Metadata: map[string]any{"log": "enabled"},
		}
		require.NoError(t, policy.MustCompile())
	})

	t.Run("it should fail to compile policy if there is an error", func(t *testing.T) {
		policy := Policy{
			Name:     "test",
			Policy:   "1 + \"Hi\"",
			Metadata: map[string]any{"log": "enabled"},
		}
		require.Error(t, policy.MustCompile())
	})
}
