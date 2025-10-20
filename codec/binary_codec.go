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

func WriteBasicType[T BasicType](buf *bytes.Buffer, v T) error {
	return binary.Write(buf, binary.BigEndian, &v)
}

func WriteBasicTypeLE[T BasicType](buf *bytes.Buffer, v T) error {
	return binary.Write(buf, binary.LittleEndian, &v)
}

// ReadBasicType reads a basic type value from the buffer using the specified byte order.
func ReadBasicType[T BasicType](buf *bytes.Buffer) (T, error) {
	var v T
	err := binary.Read(buf, binary.BigEndian, &v)
	return v, err
}

func ReadBasicTypeLE[T BasicType](buf *bytes.Buffer) (T, error) {
	var v T
	err := binary.Read(buf, binary.LittleEndian, &v)
	return v, err
}

func WriteBasicTypeList[T constraints.Unsigned, K BasicType](buf *bytes.Buffer, values []K) error {
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}
	for _, s := range values {
		if err := WriteBasicType(buf, s); err != nil {
			return err
		}
	}
	return nil
}

func WriteBasicTypeListLE[T constraints.Unsigned, K BasicType](buf *bytes.Buffer, values []K) error {
	if err := binary.Write(buf, binary.LittleEndian, T(len(values))); err != nil {
		return err
	}
	for _, s := range values {
		if err := WriteBasicType(buf, s); err != nil {
			return err
		}
	}
	return nil
}

// ReadBasicType reads a basic type value from the buffer using the specified byte order.
func ReadBasicTypeList[T constraints.Unsigned, K BasicType](buf *bytes.Buffer) ([]K, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]K, 0, count)
	var err error
	for i := 0; i < count; i++ {
		v, e := ReadBasicType[K](buf)
		if e != nil {
			return nil, e
		}
		result = append(result, v)
	}
	return result, err
}

func ReadBasicTypeListLE[T constraints.Unsigned, K BasicType](buf *bytes.Buffer) ([]K, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]K, 0, count)
	var err error
	for i := 0; i < count; i++ {
		v, e := ReadBasicType[K](buf)
		if e != nil {
			return nil, e
		}
		result = append(result, v)
	}
	return result, err
}

// ----------------------------
// String Functions
// ----------------------------

func WriteString[T constraints.Unsigned](buf *bytes.Buffer, s string) error {
	if err := binary.Write(buf, binary.BigEndian, T(len(s))); err != nil {
		return err
	}
	if _, err := buf.WriteString(s); err != nil {
		return err
	}
	return nil
}

func WriteStringLE[T constraints.Unsigned](buf *bytes.Buffer, s string) error {
	if err := binary.Write(buf, binary.LittleEndian, T(len(s))); err != nil {
		return err
	}
	if _, err := buf.WriteString(s); err != nil {
		return err
	}
	return nil
}

func ReadString[T constraints.Unsigned](buf *bytes.Buffer) (string, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return "", err
	}
	length := int(t)

	strBytes := make([]byte, length)
	_, err := io.ReadFull(buf, strBytes)
	return string(strBytes), err
}

func ReadStringLE[T constraints.Unsigned](buf *bytes.Buffer) (string, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return "", err
	}
	length := int(t)

	strBytes := make([]byte, length)
	_, err := io.ReadFull(buf, strBytes)
	return string(strBytes), err
}

func WriteFixedString(buf *bytes.Buffer, s string, fixedLen int) error {
	return WriteFixedStringWithPadding(buf, s, fixedLen, ' ', false)
}

func WriteFixedStringWithPadding(buf *bytes.Buffer, s string, fixedLen int, padding rune, fromLeft bool) error {
	data := []byte(s)
	if len(data) > fixedLen {
		if _, err := buf.Write(data[:fixedLen]); err != nil {
			return err
		}
	} else {
		if fromLeft {
			if err := Padding(buf, fixedLen-len(data), padding); err != nil {
				return err
			}
		}
		if _, err := buf.Write(data); err != nil {
			return err
		}
		if !fromLeft {
			if err := Padding(buf, fixedLen-len(data), padding); err != nil {
				return err
			}
		}
	}
	return nil
}

func Padding(buf *bytes.Buffer, paddedLen int, padding rune) error {
	paddingBytes := bytes.Repeat([]byte{byte(padding)}, paddedLen)
	if _, err := buf.Write(paddingBytes); err != nil {
		return err
	}
	return nil
}

func WriteFixedStringList[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int) error {
	return WriteFixedStringListWithPadding[T](buf, values, fixedLen, ' ', false)
}

func WriteFixedStringListWithPadding[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int, padding rune, fromLeft bool) error {
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		err := WriteFixedStringWithPadding(buf, s, fixedLen, padding, fromLeft)
		if err != nil {
			return nil
		}
	}
	return nil
}

func WriteFixedStringListLE[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int) error {
	return WriteFixedStringListWithPaddingLE[T](buf, values, fixedLen, ' ', false)
}
func WriteFixedStringListWithPaddingLE[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int, padding rune, fromLeft bool) error {
	if err := binary.Write(buf, binary.LittleEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		err := WriteFixedStringWithPadding(buf, s, fixedLen, padding, fromLeft)
		if err != nil {
			return nil
		}
	}
	return nil
}

func ReadFixedString(buf *bytes.Buffer, fixedLen int) (string, error) {
	return ReadFixedStringTrimPadding(buf, fixedLen, ' ', false)
}
func ReadFixedStringTrimPadding(buf *bytes.Buffer, fixedLen int, padding rune, fromLeft bool) (string, error) {
	strBytes := make([]byte, fixedLen)
	_, err := io.ReadFull(buf, strBytes)
	if fromLeft {
		return string(bytes.TrimLeft(strBytes, string(padding))), err
	}
	return string(bytes.TrimRight(strBytes, string(padding))), err
}

func ReadFixedStringList[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int) ([]string, error) {
	return ReadFixedStringListTrimPadding[T](buf, fixedLen, ' ', false)
}
func ReadFixedStringListTrimPadding[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int, padding rune, fromLeft bool) ([]string, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]string, 0, count)
	var err error
	for i := 0; i < count; i++ {
		str, e := ReadFixedStringTrimPadding(buf, fixedLen, padding, fromLeft)
		if e != nil {
			return nil, e
		}
		result = append(result, str)
	}
	return result, err
}

func ReadFixedStringListLE[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int) ([]string, error) {
	return ReadFixedStringListTrimPaddingLE[T](buf, fixedLen, ' ', false)
}

func ReadFixedStringListTrimPaddingLE[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int, padding rune, fromLeft bool) ([]string, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]string, 0, count)
	var err error
	for i := 0; i < count; i++ {
		str, e := ReadFixedStringTrimPadding(buf, fixedLen, padding, fromLeft)
		if e != nil {
			return nil, e
		}
		result = append(result, str)
	}
	return result, err
}

// WriteStringListLE encodes a list of strings into the buffer using little-endian format.
// T: type used for the list length prefix (e.g., uint8, uint16, uint32)
// K: type used for each string's length prefix (e.g., uint8, uint16, uint32)
func WriteStringListLE[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer, values []string) error {
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

func WriteStringList[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer, values []string) error {
	// Write the list length prefix
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		if err := binary.Write(buf, binary.BigEndian, K(len(s))); err != nil {
			return err
		}
		buf.WriteString(s)
	}
	return nil
}

// ReadStringListLE decodes a list of strings from a buffer using little-endian encoding.
// T: type used for the list length prefix (e.g., uint8, uint16, uint32)
// K: type used for each string's length prefix (e.g., uint8, uint16, uint32)
func ReadStringListLE[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer) ([]string, error) {
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

func ReadStringList[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer) ([]string, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]string, 0, count)
	for i := 0; i < count; i++ {
		var k K
		if err := binary.Read(buf, binary.BigEndian, &k); err != nil {
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

// Object
func WriteObjectList[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, values []K) error {
	// Write the list length prefix
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		if e := s.Encode(buf); e != nil {
			return e
		}
	}
	return nil
}

func ReadObjectList[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, newFn func() K) ([]K, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]K, 0, count)
	for i := 0; i < count; i++ {
		k := newFn()
		if e := k.Decode(buf); e != nil {
			return result, e
		}
		result = append(result, k)
	}
	return result, nil
}

// Object
func WriteObjectListLE[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, values []K) error {
	// Write the list length prefix
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		if e := s.Encode(buf); e != nil {
			return e
		}
	}
	return nil
}

func ReadObjectListLE[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, newFn func() K) ([]K, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]K, 0, count)
	for i := 0; i < count; i++ {
		k := newFn()
		if e := k.Decode(buf); e != nil {
			return result, e
		}
		result = append(result, k)
	}
	return result, nil
}
