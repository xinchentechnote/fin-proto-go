// Code generated by fin-protoc. DO NOT EDIT.
package bjse_trade_bin_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	msg "github.com/xinchentechnote/fin-proto-go/bjse-trade-bin/messages"
)

func TestExtendNewOrder043Codec(t *testing.T) {

	original := &msg.ExtendNewOrder043{}
	var buf bytes.Buffer
	assert.NoError(t, original.Encode(&buf))
	var decoded msg.ExtendNewOrder043
	assert.NoError(t, decoded.Decode(&buf))
	assert.Equal(t, original, &decoded)
}
