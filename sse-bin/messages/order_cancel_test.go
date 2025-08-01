// Code generated by fin-protoc. DO NOT EDIT.
package sse_bin_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	msg "github.com/xinchentechnote/fin-proto-go/sse-bin/messages"
)

func TestOrderCancelCodec(t *testing.T) {

	original := &msg.OrderCancel{
		BizId:        4,
		BizPbu:       "xxxxxxxx",
		ClOrdId:      "xxxxxxxxxx",
		SecurityId:   "xxxxxxxxxxxx",
		Account:      "xxxxxxxxxxxxx",
		OwnerType:    1,
		Side:         "x",
		OrigClOrdId:  "xxxxxxxxxx",
		TransactTime: 8,
		BranchId:     "xxxxxxxx",
		UserInfo:     "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	var buf bytes.Buffer
	assert.NoError(t, original.Encode(&buf))
	var decoded msg.OrderCancel
	assert.NoError(t, decoded.Decode(&buf))
	assert.Equal(t, original, &decoded)
}
