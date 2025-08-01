// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// Extend203715 represents the packet structure.
type Extend203715 struct {
	CashMargin string `json:"CashMargin"`
}

// NewExtend203715 creates a new instance of Extend203715.
func NewExtend203715() *Extend203715 {
	return &Extend203715{}
}

// String returns a string representation of the packet.
func (p *Extend203715) String() string {
	return fmt.Sprintf("Extend203715{CashMargin: %v}", p.CashMargin)
}

// Encode encodes the packet into a byte slice.
func (p *Extend203715) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutFixedString(buf, p.CashMargin, 1); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *Extend203715) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.CashMargin = val
	}
	return nil
}
