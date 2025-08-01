// Code generated by fin-protoc. DO NOT EDIT.
package szse_bin

import (
	"bytes"
	"fmt"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// Extend102701 represents the packet structure.
type Extend102701 struct {
	DisposalPbu       string `json:"DisposalPBU"`
	DisposalAccountId string `json:"DisposalAccountID"`
}

// NewExtend102701 creates a new instance of Extend102701.
func NewExtend102701() *Extend102701 {
	return &Extend102701{}
}

// String returns a string representation of the packet.
func (p *Extend102701) String() string {
	return fmt.Sprintf("Extend102701{DisposalPbu: %v, DisposalAccountId: %v}", p.DisposalPbu, p.DisposalAccountId)
}

// Encode encodes the packet into a byte slice.
func (p *Extend102701) Encode(buf *bytes.Buffer) error {
	// Implement encoding logic here.
	if err := codec.PutFixedString(buf, p.DisposalPbu, 6); err != nil {
		return err
	}
	if err := codec.PutFixedString(buf, p.DisposalAccountId, 12); err != nil {
		return err
	}
	return nil
}

// Decode decodes the packet from a byte slice.
func (p *Extend102701) Decode(buf *bytes.Buffer) error {
	if val, err := codec.GetFixedString(buf, 6); err != nil {
		return err
	} else {
		p.DisposalPbu = val
	}
	if val, err := codec.GetFixedString(buf, 12); err != nil {
		return err
	} else {
		p.DisposalAccountId = val
	}
	return nil
}
