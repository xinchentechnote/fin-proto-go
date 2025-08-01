// Code generated by fin-protoc. DO NOT EDIT.
package bjse_trade_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// ExecutionReport represents the packet structure.
type ExecutionReport struct {
	PartitionNo      int32             `json:"PartitionNo"`
	ReportIndex      int64             `json:"ReportIndex"`
	ApplId           string            `json:"ApplID"`
	ReportingPbuid   string            `json:"ReportingPBUID"`
	SubmittingPbuid  string            `json:"SubmittingPBUID"`
	SecurityId       string            `json:"SecurityID"`
	SecurityIdsource string            `json:"SecurityIDSource"`
	OwnerType        uint16            `json:"OwnerType"`
	ClearingFirm     string            `json:"ClearingFirm"`
	TransactTime     int64             `json:"TransactTime"`
	UserInfo         string            `json:"UserInfo"`
	OrderId          string            `json:"OrderID"`
	ClOrdId          string            `json:"ClOrdID"`
	ExecId           string            `json:"ExecID"`
	ExecType         string            `json:"ExecType"`
	OrdStatus        string            `json:"OrdStatus"`
	LastPx           int64             `json:"LastPx"`
	LastQty          int64             `json:"LastQty"`
	LeavesQty        int64             `json:"LeavesQty"`
	CumQty           int64             `json:"CumQty"`
	Side             string            `json:"Side"`
	AccountId        string            `json:"AccountID"`
	BranchId         string            `json:"BranchID"`
	ApplExtend       codec.BinaryCodec `json:"ApplExtend"`
}

// NewExecutionReport creates a new instance of ExecutionReport.
func NewExecutionReport() *ExecutionReport {
	return &ExecutionReport{}
}

// String returns a string representation of the packet.
func (p *ExecutionReport) String() string {
	return fmt.Sprintf("ExecutionReport{PartitionNo: %v, ReportIndex: %v, ApplId: %v, ReportingPbuid: %v, SubmittingPbuid: %v, SecurityId: %v, SecurityIdsource: %v, OwnerType: %v, ClearingFirm: %v, TransactTime: %v, UserInfo: %v, OrderId: %v, ClOrdId: %v, ExecId: %v, ExecType: %v, OrdStatus: %v, LastPx: %v, LastQty: %v, LeavesQty: %v, CumQty: %v, Side: %v, AccountId: %v, BranchId: %v, ApplExtend: %v}", p.PartitionNo, p.ReportIndex, p.ApplId, p.ReportingPbuid, p.SubmittingPbuid, p.SecurityId, p.SecurityIdsource, p.OwnerType, p.ClearingFirm, p.TransactTime, p.UserInfo, p.OrderId, p.ClOrdId, p.ExecId, p.ExecType, p.OrdStatus, p.LastPx, p.LastQty, p.LeavesQty, p.CumQty, p.Side, p.AccountId, p.BranchId, p.ApplExtend)
}

// Encode encodes the packet into a byte slice.
func (p *ExecutionReport) Encode(buf *bytes.Buffer) error {
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
	if err := codec.PutFixedString(buf, p.OrderId, 16); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ClOrdId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ExecId, 16); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.ExecType, 1); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.OrdStatus, 1); err != nil {
		return err
	}
	if err := codec.PutBasicTypeLE(buf, p.LastPx); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "LastPx", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.LastQty); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "LastQty", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.LeavesQty); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "LeavesQty", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.CumQty); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "CumQty", err)
	}
	if err := codec.PutFixedString(buf, p.Side, 1); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.AccountId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.BranchId, 2); err != nil {
		return err
	}
	if err := p.ApplExtend.Encode(buf); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *ExecutionReport) Decode(buf *bytes.Buffer) error {
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
	if val, err := codec.GetFixedString(buf, 16); err != nil {
		return err
	} else {
		p.OrderId = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.ClOrdId = val
	}
	if val, err := codec.GetFixedString(buf, 16); err != nil {
		return err
	} else {
		p.ExecId = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.ExecType = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.OrdStatus = val
	}
	if val, err := codec.GetBasicTypeLE[int64](buf); err != nil {
		return err
	} else {
		p.LastPx = val
	}
	if val, err := codec.GetBasicTypeLE[int64](buf); err != nil {
		return err
	} else {
		p.LastQty = val
	}
	if val, err := codec.GetBasicTypeLE[int64](buf); err != nil {
		return err
	} else {
		p.LeavesQty = val
	}
	if val, err := codec.GetBasicTypeLE[int64](buf); err != nil {
		return err
	} else {
		p.CumQty = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.Side = val
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
	switch p.ApplId {
	case "010":
		p.ApplExtend = &ReportExtend010{}
	case "040":
		p.ApplExtend = &ReportExtend040{}
	case "050":
		p.ApplExtend = &ReportExtend050{}
	default:
		return fmt.Errorf("unsupported ApplId: %v", p.ApplId)
	}
	if err := p.ApplExtend.Decode(buf); err != nil {
		return err
	}
	return nil
}
