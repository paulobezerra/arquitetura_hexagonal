package handler

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandler_jsonError(t *testing.T) {
	msg := "Hello Test"
	result := jsonError(msg)

	require.Equal(t, []byte(`{"message":"Hello Test"}`), result)
}
