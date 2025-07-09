package codec

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"

	"golang.org/x/exp/constraints"
)

type BinaryCodec interface {
	Encode(buf *bytes.Buffer) error
	Decode(buf *bytes.Buffer) error
}

type BasicType interface {
	~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// ----------------------------
// String Functions
// ----------------------------

func PutString[T constraints.Unsigned](buf *bytes.Buffer, s string) error {
	if err := binary.Write(buf, binary.BigEndian, T(len(s))); err != nil {
		return err
	}
	if _, err := buf.WriteString(s); err != nil {
		return err
	}
	return nil
}

func PutStringLE[T constraints.Unsigned](buf *bytes.Buffer, s string) error {
	if err := binary.Write(buf, binary.LittleEndian, T(len(s))); err != nil {
		return err
	}
	if _, err := buf.WriteString(s); err != nil {
		return err
	}
	return nil
}

func GetString[T constraints.Unsigned](buf *bytes.Buffer) (string, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return "", err
	}
	length := int(t)

	strBytes := make([]byte, length)
	_, err := io.ReadFull(buf, strBytes)
	return string(strBytes), err
}

func GetStringLE[T constraints.Unsigned](buf *bytes.Buffer) (string, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return "", err
	}
	length := int(t)

	strBytes := make([]byte, length)
	_, err := io.ReadFull(buf, strBytes)
	return string(strBytes), err
}

func PutFixedString(buf *bytes.Buffer, s string, fixedLen int) error {
	data := []byte(s)
	if len(data) > fixedLen {
		if _, err := buf.Write(data[:fixedLen]); err != nil {
			return err
		}
	} else {
		if _, err := buf.Write(data); err != nil {
			return err
		}
		if _, err := buf.Write(bytes.Repeat([]byte{0}, fixedLen-len(data))); err != nil {
			return err
		}
	}
	return nil
}

func GetFixedString(buf *bytes.Buffer, fixedLen int) (string, error) {
	strBytes := make([]byte, fixedLen)
	_, err := io.ReadFull(buf, strBytes)
	return string(bytes.TrimRight(strBytes, "\x00")), err
}

// PutStringListLE encodes a list of strings into the buffer using little-endian format.
// T: type used for the list length prefix (e.g., uint8, uint16, uint32)
// K: type used for each string's length prefix (e.g., uint8, uint16, uint32)
func PutStringListLE[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer, values []string) error {
	// Write the list length prefix
	if err := binary.Write(buf, binary.LittleEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		if err := binary.Write(buf, binary.LittleEndian, K(len(s))); err != nil {
			return err
		}
		buf.WriteString(s)
	}
	return nil
}

// GetStringListLE decodes a list of strings from a buffer using little-endian encoding.
// T: type used for the list length prefix (e.g., uint8, uint16, uint32)
// K: type used for each string's length prefix (e.g., uint8, uint16, uint32)
func GetStringListLE[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer) ([]string, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]string, 0, count)
	for i := 0; i < count; i++ {
		var k K
		if err := binary.Read(buf, binary.LittleEndian, &k); err != nil {
			return nil, err
		}
		length := int(k)

		strBytes := make([]byte, length)
		n, err := buf.Read(strBytes)
		if err != nil || n != length {
			return nil, errors.New("incomplete string bytes")
		}

		result = append(result, string(strBytes))
	}
	return result, nil
}

// GetBasicType reads a basic type value from the buffer using the specified byte order.
func GetBasicType[T BasicType](buf *bytes.Buffer, order binary.ByteOrder) (T, error) {
	var v T
	err := binary.Read(buf, order, &v)
	return v, err
}
