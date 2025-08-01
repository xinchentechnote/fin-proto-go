// Code generated by fin-protoc. DO NOT EDIT.
package sse_bin_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	msg "github.com/xinchentechnote/fin-proto-go/sse-bin/messages"
)

func TestOrderRejectCodec(t *testing.T) {

	original := &msg.OrderReject{
		BizId:        4,
		BizPbu:       "xxxxxxxx",
		ClOrdId:      "xxxxxxxxxx",
		SecurityId:   "xxxxxxxxxxxx",
		OrdRejReason: 4,
		TradeDate:    4,
		TransactTime: 8,
		UserInfo:     "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	var buf bytes.Buffer
	assert.NoError(t, original.Encode(&buf))
	var decoded msg.OrderReject
	assert.NoError(t, decoded.Decode(&buf))
	assert.Equal(t, original, &decoded)
}
