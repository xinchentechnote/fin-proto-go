package sample_test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	sample "github.com/xinchentechnote/fin-proto-go/sample/messages"
)

func TestRiskControlRequestCodec(t *testing.T) {
	original := sample.RiskControlRequest{
		UniqueOrderID: "example",
		ClOrdID:       "aaaaaaaaaaaaaaaa",
		MarketID:      "aaa",
		SecurityID:    "aaaaaaaaaaaa",
		Side:          'a',
		OrderType:     'a',
		Price:         123456789,
		Qty:           123456,
		ExtraInfo:     []string{"example", "test"},
		SubOrder: sample.SubOrder{
			ClOrdID: "aaaaaaaaaaaaaaaa",
			Price:   123456789,
			Qty:     123456,
		},
	}

	var buf bytes.Buffer
	assert.NoError(t, original.Encode(&buf))
	reader := bytes.NewReader(buf.Bytes())
	var decoded sample.RiskControlRequest
	assert.NoError(t, decoded.Decode(reader))
	assert.Equal(t, original, decoded)
}
