// Code generated by fin-protoc. DO NOT EDIT.
package bjse_trade_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// CancelReject represents the packet structure.
type CancelReject struct {
	PartitionNo      int32  `json:"PartitionNo"`
	ReportIndex      int64  `json:"ReportIndex"`
	ApplId           string `json:"ApplID"`
	ReportingPbuid   string `json:"ReportingPBUID"`
	SubmittingPbuid  string `json:"SubmittingPBUID"`
	SecurityId       string `json:"SecurityID"`
	SecurityIdsource string `json:"SecurityIDSource"`
	OwnerType        uint16 `json:"OwnerType"`
	ClearingFirm     string `json:"ClearingFirm"`
	TransactTime     int64  `json:"TransactTime"`
	UserInfo         string `json:"UserInfo"`
	ClOrdId          string `json:"ClOrdID"`
	OrigClOrdId      string `json:"OrigClOrdID"`
	AccountId        string `json:"AccountID"`
	BranchId         string `json:"BranchID"`
	OrdStatus        string `json:"OrdStatus"`
	CxlRejReason     uint16 `json:"CxlRejReason"`
	RejectText       string `json:"RejectText"`
	OrderId          string `json:"OrderID"`
}

// NewCancelReject creates a new instance of CancelReject.
func NewCancelReject() *CancelReject {
	return &CancelReject{}
}

// String returns a string representation of the packet.
func (p *CancelReject) String() string {
	return fmt.Sprintf("CancelReject{PartitionNo: %v, ReportIndex: %v, ApplId: %v, ReportingPbuid: %v, SubmittingPbuid: %v, SecurityId: %v, SecurityIdsource: %v, OwnerType: %v, ClearingFirm: %v, TransactTime: %v, UserInfo: %v, ClOrdId: %v, OrigClOrdId: %v, AccountId: %v, BranchId: %v, OrdStatus: %v, CxlRejReason: %v, RejectText: %v, OrderId: %v}", p.PartitionNo, p.ReportIndex, p.ApplId, p.ReportingPbuid, p.SubmittingPbuid, p.SecurityId, p.SecurityIdsource, p.OwnerType, p.ClearingFirm, p.TransactTime, p.UserInfo, p.ClOrdId, p.OrigClOrdId, p.AccountId, p.BranchId, p.OrdStatus, p.CxlRejReason, p.RejectText, p.OrderId)
}

// Encode encodes the packet into a byte slice.
func (p *CancelReject) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutBasicTypeLE(buf, p.PartitionNo); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "PartitionNo", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.ReportIndex); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "ReportIndex", err)
	}
	if err := codec.PutFixedString(buf, p.ApplId, 3); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ReportingPbuid, 6); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.SubmittingPbuid, 6); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.SecurityId, 8); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.SecurityIdsource, 4); err != nil {
		return err
	}
	if err := codec.PutBasicTypeLE(buf, p.OwnerType); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "OwnerType", err)
	}
	if err := codec.PutFixedString(buf, p.ClearingFirm, 2); err != nil {
		return err
	}
	if err := codec.PutBasicTypeLE(buf, p.TransactTime); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TransactTime", err)
	}
	if err := codec.PutFixedString(buf, p.UserInfo, 32); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ClOrdId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.OrigClOrdId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.AccountId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.BranchId, 2); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.OrdStatus, 1); err != nil {
		return err
	}
	if err := codec.PutBasicTypeLE(buf, p.CxlRejReason); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "CxlRejReason", err)
	}
	if err := codec.PutFixedString(buf, p.RejectText, 16); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.OrderId, 16); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *CancelReject) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetBasicTypeLE[int32](buf); err != nil {
		return err
	} else {
		p.PartitionNo = val
	}
	if val, err := codec.GetBasicTypeLE[int64](buf); err != nil {
		return err
	} else {
		p.ReportIndex = val
	}
	if val, err := codec.GetFixedString(buf, 3); err != nil {
		return err
	} else {
		p.ApplId = val
	}
	if val, err := codec.GetFixedString(buf, 6); err != nil {
		return err
	} else {
		p.ReportingPbuid = val
	}
	if val, err := codec.GetFixedString(buf, 6); err != nil {
		return err
	} else {
		p.SubmittingPbuid = val
	}
	if val, err := codec.GetFixedString(buf, 8); err != nil {
		return err
	} else {
		p.SecurityId = val
	}
	if val, err := codec.GetFixedString(buf, 4); err != nil {
		return err
	} else {
		p.SecurityIdsource = val
	}
	if val, err := codec.GetBasicTypeLE[uint16](buf); err != nil {
		return err
	} else {
		p.OwnerType = val
	}
	if val, err := codec.GetFixedString(buf, 2); err != nil {
		return err
	} else {
		p.ClearingFirm = val
	}
	if val, err := codec.GetBasicTypeLE[int64](buf); err != nil {
		return err
	} else {
		p.TransactTime = val
	}
	if val, err := codec.GetFixedString(buf, 32); err != nil {
		return err
	} else {
		p.UserInfo = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.ClOrdId = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.OrigClOrdId = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.AccountId = val
	}
	if val, err := codec.GetFixedString(buf, 2); err != nil {
		return err
	} else {
		p.BranchId = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.OrdStatus = val
	}
	if val, err := codec.GetBasicTypeLE[uint16](buf); err != nil {
		return err
	} else {
		p.CxlRejReason = val
	}
	if val, err := codec.GetFixedString(buf, 16); err != nil {
		return err
	} else {
		p.RejectText = val
	}
	if val, err := codec.GetFixedString(buf, 16); err != nil {
		return err
	} else {
		p.OrderId = val
	}
	return nil
}
