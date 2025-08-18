package szse_bin

import (
	"bytes"
	"testing"

	szse_bin "github.com/xinchentechnote/fin-proto-go/szse-bin/messages"
)

func sampleNewOrder() *szse_bin.NewOrder {
	return &szse_bin.NewOrder{
		ApplId:            "010",
		SubmittingPbuid:   "b001",
		SecurityId:        "000001",
		SecurityIdsource:  "101",
		OwnerType:         1,
		ClearingFirm:      "c01",
		TransactTime:      20250101123000,
		UserInfo:          "u01",
		ClOrdId:           "c00001",
		AccountId:         "a01",
		BranchId:          "b01",
		OrderRestrictions: "o01",
		Side:              "1",
		OrdType:           "1",
		OrderQty:          1000,
		Price:             90,
	}
}

func BenchmarkEncodeNewOrder(b *testing.B) {
	order := sampleNewOrder()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		_ = order.Encode(&buf)
	}
}

func BenchmarkDecodeNewOrder(b *testing.B) {
	order := sampleNewOrder()
	var buf bytes.Buffer
	order.Encode(&buf)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var order szse_bin.NewOrder
		_ = order.Decode(&buf)
	}
}
