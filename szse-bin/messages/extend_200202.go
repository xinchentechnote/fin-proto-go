// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// Extend200202 represents the packet structure.
type Extend200202 struct {
	StopPx         int64  `json:"StopPx"`
	MinQty         int64  `json:"MinQty"`
	MaxPriceLevels uint16 `json:"MaxPriceLevels"`
	TimeInForce    string `json:"TimeInForce"`
}

// NewExtend200202 creates a new instance of Extend200202.
func NewExtend200202() *Extend200202 {
	return &Extend200202{}
}

// String returns a string representation of the packet.
func (p *Extend200202) String() string {
	return fmt.Sprintf("Extend200202{StopPx: %v, MinQty: %v, MaxPriceLevels: %v, TimeInForce: %v}", p.StopPx, p.MinQty, p.MaxPriceLevels, p.TimeInForce)
}

// Encode encodes the packet into a byte slice.
func (p *Extend200202) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutBasicType(buf, p.StopPx); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "StopPx", err)
	}
	if err := codec.PutBasicType(buf, p.MinQty); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "MinQty", err)
	}
	if err := codec.PutBasicType(buf, p.MaxPriceLevels); err != nil {
		return fmt.Errorf("failed to encode %s: %w", "MaxPriceLevels", err)
	}
	if err := codec.PutFixedString(buf, p.TimeInForce, 1); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *Extend200202) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetBasicType[int64](buf); err != nil {
		return err
	} else {
		p.StopPx = val
	}
	if val, err := codec.GetBasicType[int64](buf); err != nil {
		return err
	} else {
		p.MinQty = val
	}
	if val, err := codec.GetBasicType[uint16](buf); err != nil {
		return err
	} else {
		p.MaxPriceLevels = val
	}
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.TimeInForce = val
	}
	return nil
}
