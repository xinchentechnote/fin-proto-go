package sample_bin

import (
	"bytes"
	"encoding/binary"

	"github.com/xinchentechnote/fin-proto-go/internal/codec"
)

type SubOrder struct {
	ClOrdID string
	Price   uint64
	Qty     uint32
}

func (s *SubOrder) Encode(buf *bytes.Buffer) {
	codec.PutFixedString(buf, s.ClOrdID, 16)
	binary.Write(buf, binary.BigEndian, s.Price)
	binary.Write(buf, binary.BigEndian, s.Qty)
}

func (s *SubOrder) Decode(buf *bytes.Buffer) error {
	var err error
	if s.ClOrdID, err = codec.GetFixedString(buf, 16); err != nil {
		return err
	}
	if s.Price, err = codec.GetBasicType[uint64](buf); err != nil {
		return err
	}
	if s.Qty, err = codec.GetBasicType[uint32](buf); err != nil {
		return err
	}
	return nil
}

type RiskControlRequest struct {
	UniqueOrderID string
	ClOrdID       string
	MarketID      string
	SecurityID    string
	Side          byte
	OrderType     byte
	Price         uint64
	Qty           uint32
	ExtraInfo     []string
	SubOrder      SubOrder
}

func (r *RiskControlRequest) Encode(buf *bytes.Buffer) error {
	codec.PutString[uint16](buf, r.UniqueOrderID)
	codec.PutFixedString(buf, r.ClOrdID, 16)
	codec.PutFixedString(buf, r.MarketID, 3)
	codec.PutFixedString(buf, r.SecurityID, 12)
	binary.Write(buf, binary.BigEndian, r.Side)
	binary.Write(buf, binary.BigEndian, r.OrderType)
	binary.Write(buf, binary.BigEndian, r.Price)
	binary.Write(buf, binary.BigEndian, r.Qty)
	codec.PutStringList[uint16, uint16](buf, r.ExtraInfo)
	r.SubOrder.Encode(buf)
	return nil
}

func (r *RiskControlRequest) Decode(buf *bytes.Buffer) error {
	var err error
	if r.UniqueOrderID, err = codec.GetString[uint16](buf); err != nil {
		return err
	}
	if r.ClOrdID, err = codec.GetFixedString(buf, 16); err != nil {
		return err
	}
	if r.MarketID, err = codec.GetFixedString(buf, 3); err != nil {
		return err
	}
	if r.SecurityID, err = codec.GetFixedString(buf, 12); err != nil {
		return err
	}
	if r.Side, err = codec.GetBasicType[byte](buf); err != nil {
		return err
	}
	if r.OrderType, err = codec.GetBasicType[byte](buf); err != nil {
		return err
	}
	if r.Price, err = codec.GetBasicType[uint64](buf); err != nil {
		return err
	}
	if r.Qty, err = codec.GetBasicType[uint32](buf); err != nil {
		return err
	}
	if r.ExtraInfo, err = codec.GetStringList[uint16, uint16](buf); err != nil {
		return err
	}
	if err = r.SubOrder.Decode(buf); err != nil {
		return err
	}
	return nil
}
