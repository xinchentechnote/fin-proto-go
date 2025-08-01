// Code generated by fin-protoc. DO NOT EDIT.
package sse_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// NewOrderSingle represents the packet structure.
type NewOrderSingle struct {
	BizId        uint32 `json:"BizID"`
	BizPbu       string `json:"BizPbu"`
	ClOrdId      string `json:"ClOrdID"`
	SecurityId   string `json:"SecurityID"`
	Account      string `json:"Account"`
	OwnerType    uint8  `json:"OwnerType"`
	Side         string `json:"Side"`
	Price        int64  `json:"Price"`
	OrderQty     int64  `json:"OrderQty"`
	OrdType      string `json:"OrdType"`
	TimeInForce  string `json:"TimeInForce"`
	TransactTime uint64 `json:"TransactTime"`
	CreditTag    string `json:"CreditTag"`
	ClearingFirm string `json:"ClearingFirm"`
	BranchId     string `json:"BranchID"`
	UserInfo     string `json:"UserInfo"`
}

// NewNewOrderSingle creates a new instance of NewOrderSingle.
func NewNewOrderSingle() *NewOrderSingle {
	return &NewOrderSingle{}
}

// String returns a string representation of the packet.
func (p *NewOrderSingle) String() string {
	return fmt.Sprintf("NewOrderSingle{BizId: %v, BizPbu: %v, ClOrdId: %v, SecurityId: %v, Account: %v, OwnerType: %v, Side: %v, Price: %v, OrderQty: %v, OrdType: %v, TimeInForce: %v, TransactTime: %v, CreditTag: %v, ClearingFirm: %v, BranchId: %v, UserInfo: %v}", p.BizId, p.BizPbu, p.ClOrdId, p.SecurityId, p.Account, p.OwnerType, p.Side, p.Price, p.OrderQty, p.OrdType, p.TimeInForce, p.TransactTime, p.CreditTag, p.ClearingFirm, p.BranchId, p.UserInfo)
}

// Encode encodes the packet into a byte slice.
func (p *NewOrderSingle) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutBasicType(buf, p.BizId); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "BizID", err)
	}
	if err := codec.PutFixedString(buf, p.BizPbu, 8); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ClOrdId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.SecurityId, 12); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.Account, 13); err != nil {
		return err
	}
	if err := codec.PutBasicType(buf, p.OwnerType); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "OwnerType", err)
	}
	if err := codec.PutFixedString(buf, p.Side, 1); err != nil {
		return err
	}
	if err := codec.PutBasicType(buf, p.Price); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "Price", err)
	}
	if err := codec.PutBasicType(buf, p.OrderQty); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "OrderQty", err)
	}
	if err := codec.PutFixedString(buf, p.OrdType, 1); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.TimeInForce, 1); err != nil {
		return err
	}
	if err := codec.PutBasicType(buf, p.TransactTime); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TransactTime", err)
	}
	if err := codec.PutFixedString(buf, p.CreditTag, 2); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ClearingFirm, 8); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.BranchId, 8); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.UserInfo, 32); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *NewOrderSingle) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetBasicType[uint32](buf); err != nil {
		return err
	} else {
		p.BizId = val
	}
	if val, err := codec.GetFixedString(buf, 8); err != nil {
		return err
	} else {
		p.BizPbu = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.ClOrdId = val
	}
	if val, err := codec.GetFixedString(buf, 12); err != nil {
		return err
	} else {
		p.SecurityId = val
	}
	if val, err := codec.GetFixedString(buf, 13); err != nil {
		return err
	} else {
		p.Account = val
	}
	if val, err := codec.GetBasicType[uint8](buf); err != nil {
		return err
	} else {
		p.OwnerType = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.Side = val
	}
	if val, err := codec.GetBasicType[int64](buf); err != nil {
		return err
	} else {
		p.Price = val
	}
	if val, err := codec.GetBasicType[int64](buf); err != nil {
		return err
	} else {
		p.OrderQty = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.OrdType = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.TimeInForce = val
	}
	if val, err := codec.GetBasicType[uint64](buf); err != nil {
		return err
	} else {
		p.TransactTime = val
	}
	if val, err := codec.GetFixedString(buf, 2); err != nil {
		return err
	} else {
		p.CreditTag = val
	}
	if val, err := codec.GetFixedString(buf, 8); err != nil {
		return err
	} else {
		p.ClearingFirm = val
	}
	if val, err := codec.GetFixedString(buf, 8); err != nil {
		return err
	} else {
		p.BranchId = val
	}
	if val, err := codec.GetFixedString(buf, 32); err != nil {
		return err
	} else {
		p.UserInfo = val
	}
	return nil
}
