package sse_bin

import (
	"bytes"
	"testing"

	sse_bin "github.com/xinchentechnote/fin-proto-go/sse-bin/messages"
)

func newSimpleOrder() *sse_bin.NewOrderSingle {
	return &sse_bin.NewOrderSingle{
		BizId:        123456,
		BizPbu:       "010",
		ClOrdId:      "c00001",
		SecurityId:   "000001",
		Account:      "a01",
		OwnerType:    1,
		Side:         "1",
		Price:        90,
		OrderQty:     1000,
		OrdType:      "1",
		TimeInForce:  "1",
		TransactTime: 20250101123000,
		CreditTag:    "01",
		ClearingFirm: "c01",
		BranchId:     "b01",
		UserInfo:     "u01",
	}
}

func BenchmarkEncodeNewOrderSingle(b *testing.B) {
	order := newSimpleOrder()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = order.Encode(&buf)
	}
}

func BenchmarkDecodeNewOrderSingle(b *testing.B) {
	order := newSimpleOrder()
	var buf bytes.Buffer
	order.Encode(&buf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		order := &sse_bin.NewOrderSingle{}
		_ = order.Decode(&buf)
	}
}
