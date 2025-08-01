// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	msg "github.com/xinchentechnote/fin-proto-go/szse-bin/messages"
)

func TestSzseBinaryCodec(t *testing.T) {

	body := &msg.Logon{
		SenderCompId:     "xxxxxxxxxxxxxxxxxxxx",
		TargetCompId:     "xxxxxxxxxxxxxxxxxxxx",
		HeartBtint:       4,
		Password:         "xxxxxxxxxxxxxxxx",
		DefaultApplVerId: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	original := &msg.SzseBinary{
		BodyLength: 4,
		MsgType:    1,
		Body:       body,
		Checksum:   4,
	}
	var buf bytes.Buffer
	assert.NoError(t, original.Encode(&buf))
	var decoded msg.SzseBinary
	assert.NoError(t, decoded.Decode(&buf))
	assert.Equal(t, original, &decoded)
}
