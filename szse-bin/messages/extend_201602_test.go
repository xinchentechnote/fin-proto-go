// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	msg "github.com/xinchentechnote/fin-proto-go/szse-bin/messages"
)

func TestExtend201602Codec(t *testing.T) {

	original := &msg.Extend201602{
		ContractAccountCode: "xxxxxx",
	}
	var buf bytes.Buffer
	assert.NoError(t, original.Encode(&buf))
	var decoded msg.Extend201602
	assert.NoError(t, decoded.Decode(&buf))
	assert.Equal(t, original, &decoded)
}
