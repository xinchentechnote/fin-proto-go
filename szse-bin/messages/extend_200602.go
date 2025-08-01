// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// Extend200602 represents the packet structure.
type Extend200602 struct {
	CashMargin string `json:"CashMargin"`
}

// NewExtend200602 creates a new instance of Extend200602.
func NewExtend200602() *Extend200602 {
	return &Extend200602{}
}

// String returns a string representation of the packet.
func (p *Extend200602) String() string {
	return fmt.Sprintf("Extend200602{CashMargin: %v}", p.CashMargin)
}

// Encode encodes the packet into a byte slice.
func (p *Extend200602) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutFixedString(buf, p.CashMargin, 1); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *Extend200602) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetFixedString(buf, 1); err != nil {
		return err
	} else {
		p.CashMargin = val
	}
	return nil
}
