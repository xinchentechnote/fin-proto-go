// Code generated by fin-protoc. DO NOT EDIT.
package bjse_trade_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// TradeCaptureReportAck represents the packet structure.
type TradeCaptureReportAck struct {
	PartitionNo             int32             `json:"PartitionNo"`
	ReportIndex             int64             `json:"ReportIndex"`
	ApplId                  string            `json:"ApplID"`
	ReportingPbuid          string            `json:"ReportingPBUID"`
	SubmittingPbuid         string            `json:"SubmittingPBUID"`
	SecurityId              string            `json:"SecurityID"`
	SecurityIdsource        string            `json:"SecurityIDSource"`
	OwnerType               uint16            `json:"OwnerType"`
	ClearingFirm            string            `json:"ClearingFirm"`
	TransactTime            int64             `json:"TransactTime"`
	UserInfo                string            `json:"UserInfo"`
	TradeId                 string            `json:"TradeID"`
	TradeReportId           string            `json:"TradeReportID"`
	TradeReportType         uint8             `json:"TradeReportType"`
	TradeReportTransType    uint8             `json:"TradeReportTransType"`
	TradeHandlingInstr      string            `json:"TradeHandlingInstr"`
	TradeReportRefId        string            `json:"TradeReportRefID"`
	TrdAckStatus            uint8             `json:"TrdAckStatus"`
	TrdRptStatus            uint8             `json:"TrdRptStatus"`
	TradeReportRejectReason uint16            `json:"TradeReportRejectReason"`
	LastPx                  int64             `json:"LastPx"`
	LastQty                 int64             `json:"LastQty"`
	TrdType                 uint16            `json:"TrdType"`
	TrdSubType              uint16            `json:"TrdSubType"`
	ConfirmId               uint32            `json:"ConfirmID"`
	ExecId                  string            `json:"ExecID"`
	Side                    string            `json:"Side"`
	Pbuid                   string            `json:"PBUID"`
	AccountId               string            `json:"AccountID"`
	BranchId                string            `json:"BranchID"`
	CounterPartyPbuid       string            `json:"CounterPartyPBUID"`
	CounterPartyAccountId   string            `json:"CounterPartyAccountID"`
	CounterPartyBranchId    string            `json:"CounterPartyBranchID"`
	ApplExtend              codec.BinaryCodec `json:"ApplExtend"`
}

// NewTradeCaptureReportAck creates a new instance of TradeCaptureReportAck.
func NewTradeCaptureReportAck() *TradeCaptureReportAck {
	return &TradeCaptureReportAck{}
}

// String returns a string representation of the packet.
func (p *TradeCaptureReportAck) String() string {
	return fmt.Sprintf("TradeCaptureReportAck{PartitionNo: %v, ReportIndex: %v, ApplId: %v, ReportingPbuid: %v, SubmittingPbuid: %v, SecurityId: %v, SecurityIdsource: %v, OwnerType: %v, ClearingFirm: %v, TransactTime: %v, UserInfo: %v, TradeId: %v, TradeReportId: %v, TradeReportType: %v, TradeReportTransType: %v, TradeHandlingInstr: %v, TradeReportRefId: %v, TrdAckStatus: %v, TrdRptStatus: %v, TradeReportRejectReason: %v, LastPx: %v, LastQty: %v, TrdType: %v, TrdSubType: %v, ConfirmId: %v, ExecId: %v, Side: %v, Pbuid: %v, AccountId: %v, BranchId: %v, CounterPartyPbuid: %v, CounterPartyAccountId: %v, CounterPartyBranchId: %v, ApplExtend: %v}", p.PartitionNo, p.ReportIndex, p.ApplId, p.ReportingPbuid, p.SubmittingPbuid, p.SecurityId, p.SecurityIdsource, p.OwnerType, p.ClearingFirm, p.TransactTime, p.UserInfo, p.TradeId, p.TradeReportId, p.TradeReportType, p.TradeReportTransType, p.TradeHandlingInstr, p.TradeReportRefId, p.TrdAckStatus, p.TrdRptStatus, p.TradeReportRejectReason, p.LastPx, p.LastQty, p.TrdType, p.TrdSubType, p.ConfirmId, p.ExecId, p.Side, p.Pbuid, p.AccountId, p.BranchId, p.CounterPartyPbuid, p.CounterPartyAccountId, p.CounterPartyBranchId, p.ApplExtend)
}

// Encode encodes the packet into a byte slice.
func (p *TradeCaptureReportAck) Encode(buf *bytes.Buffer) error {
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
	if err := codec.PutFixedString(buf, p.TradeId, 16); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.TradeReportId, 10); err != nil {
		return err
	}
	if err := codec.PutBasicTypeLE(buf, p.TradeReportType); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TradeReportType", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.TradeReportTransType); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TradeReportTransType", err)
	}
	if err := codec.PutFixedString(buf, p.TradeHandlingInstr, 1); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.TradeReportRefId, 10); err != nil {
		return err
	}
	if err := codec.PutBasicTypeLE(buf, p.TrdAckStatus); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TrdAckStatus", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.TrdRptStatus); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TrdRptStatus", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.TradeReportRejectReason); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TradeReportRejectReason", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.LastPx); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "LastPx", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.LastQty); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "LastQty", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.TrdType); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TrdType", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.TrdSubType); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "TrdSubType", err)
	}
	if err := codec.PutBasicTypeLE(buf, p.ConfirmId); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "ConfirmID", err)
	}
	if err := codec.PutFixedString(buf, p.ExecId, 16); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.Side, 1); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.Pbuid, 6); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.AccountId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.BranchId, 2); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.CounterPartyPbuid, 6); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.CounterPartyAccountId, 10); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.CounterPartyBranchId, 2); err != nil {
		return err
	}
	if err := p.ApplExtend.Encode(buf); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *TradeCaptureReportAck) Decode(buf *bytes.Buffer) error {
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
		p.TradeId = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.TradeReportId = val
	}
	if val, err := codec.GetBasicTypeLE[uint8](buf); err != nil {
		return err
	} else {
		p.TradeReportType = val
	}
	if val, err := codec.GetBasicTypeLE[uint8](buf); err != nil {
		return err
	} else {
		p.TradeReportTransType = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.TradeHandlingInstr = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.TradeReportRefId = val
	}
	if val, err := codec.GetBasicTypeLE[uint8](buf); err != nil {
		return err
	} else {
		p.TrdAckStatus = val
	}
	if val, err := codec.GetBasicTypeLE[uint8](buf); err != nil {
		return err
	} else {
		p.TrdRptStatus = val
	}
	if val, err := codec.GetBasicTypeLE[uint16](buf); err != nil {
		return err
	} else {
		p.TradeReportRejectReason = val
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
	if val, err := codec.GetBasicTypeLE[uint16](buf); err != nil {
		return err
	} else {
		p.TrdType = val
	}
	if val, err := codec.GetBasicTypeLE[uint16](buf); err != nil {
		return err
	} else {
		p.TrdSubType = val
	}
	if val, err := codec.GetBasicTypeLE[uint32](buf); err != nil {
		return err
	} else {
		p.ConfirmId = val
	}
	if val, err := codec.GetFixedString(buf, 16); err != nil {
		return err
	} else {
		p.ExecId = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.Side = val
	}
	if val, err := codec.GetFixedString(buf, 6); err != nil {
		return err
	} else {
		p.Pbuid = val
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
	if val, err := codec.GetFixedString(buf, 6); err != nil {
		return err
	} else {
		p.CounterPartyPbuid = val
	}
	if val, err := codec.GetFixedString(buf, 10); err != nil {
		return err
	} else {
		p.CounterPartyAccountId = val
	}
	if val, err := codec.GetFixedString(buf, 2); err != nil {
		return err
	} else {
		p.CounterPartyBranchId = val
	}
	switch p.ApplId {
	case "031":
		p.ApplExtend = &TradeCaptureReportAckExtend031{}
	case "051":
		p.ApplExtend = &TradeCaptureReportAckExtend051{}
	case "060":
		p.ApplExtend = &TradeCaptureReportAckExtend060{}
	case "061":
		p.ApplExtend = &TradeCaptureReportAckExtend061{}
	case "062":
		p.ApplExtend = &TradeCaptureReportAckExtend062{}
	default:
		return fmt.Errorf("unsupported ApplId: %v", p.ApplId)
	}
	if err := p.ApplExtend.Decode(buf); err != nil {
		return err
	}
	return nil
}
