// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// Extend101501 represents the packet structure.
type Extend101501 struct {
	ShareProperty string `json:"ShareProperty"`
}

// NewExtend101501 creates a new instance of Extend101501.
func NewExtend101501() *Extend101501 {
	return &Extend101501{}
}

// String returns a string representation of the packet.
func (p *Extend101501) String() string {
	return fmt.Sprintf("Extend101501{ShareProperty: %v}", p.ShareProperty)
}

// Encode encodes the packet into a byte slice.
func (p *Extend101501) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutFixedString(buf, p.ShareProperty, 2); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *Extend101501) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetFixedString(buf, 2); err != nil {
		return err
	} else {
		p.ShareProperty = val
	}
	return nil
}
