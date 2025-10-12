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

func PutBasicType[T BasicType](buf *bytes.Buffer, v T) error {
	return binary.Write(buf, binary.BigEndian, &v)
}

func PutBasicTypeLE[T BasicType](buf *bytes.Buffer, v T) error {
	return binary.Write(buf, binary.LittleEndian, &v)
}

// GetBasicType reads a basic type value from the buffer using the specified byte order.
func GetBasicType[T BasicType](buf *bytes.Buffer) (T, error) {
	var v T
	err := binary.Read(buf, binary.BigEndian, &v)
	return v, err
}

func GetBasicTypeLE[T BasicType](buf *bytes.Buffer) (T, error) {
	var v T
	err := binary.Read(buf, binary.LittleEndian, &v)
	return v, err
}

func PutBasicTypeList[T constraints.Unsigned, K BasicType](buf *bytes.Buffer, values []K) error {
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}
	for _, s := range values {
		if err := PutBasicType(buf, s); err != nil {
			return err
		}
	}
	return nil
}

func PutBasicTypeListLE[T constraints.Unsigned, K BasicType](buf *bytes.Buffer, values []K) error {
	if err := binary.Write(buf, binary.LittleEndian, T(len(values))); err != nil {
		return err
	}
	for _, s := range values {
		if err := PutBasicType(buf, s); err != nil {
			return err
		}
	}
	return nil
}

// GetBasicType reads a basic type value from the buffer using the specified byte order.
func GetBasicTypeList[T constraints.Unsigned, K BasicType](buf *bytes.Buffer) ([]K, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]K, 0, count)
	var err error
	for i := 0; i < count; i++ {
		v, e := GetBasicType[K](buf)
		if e != nil {
			return nil, e
		}
		result = append(result, v)
	}
	return result, err
}

func GetBasicTypeListLE[T constraints.Unsigned, K BasicType](buf *bytes.Buffer) ([]K, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]K, 0, count)
	var err error
	for i := 0; i < count; i++ {
		v, e := GetBasicType[K](buf)
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
	return PutFixedStringWithPadding(buf, s, fixedLen, ' ', false)
}

func PutFixedStringWithPadding(buf *bytes.Buffer, s string, fixedLen int, padding rune, fromLeft bool) error {
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

func PutFixedStringList[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int) error {
	return PutFixedStringListWithPadding[T](buf, values, fixedLen, ' ', false)
}

func PutFixedStringListWithPadding[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int, padding rune, fromLeft bool) error {
	if err := binary.Write(buf, binary.BigEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		err := PutFixedStringWithPadding(buf, s, fixedLen, padding, fromLeft)
		if err != nil {
			return nil
		}
	}
	return nil
}

func PutFixedStringListLE[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int) error {
	return PutFixedStringListWithPaddingLE[T](buf, values, fixedLen, ' ', false)
}
func PutFixedStringListWithPaddingLE[T constraints.Unsigned](buf *bytes.Buffer, values []string, fixedLen int, padding rune, fromLeft bool) error {
	if err := binary.Write(buf, binary.LittleEndian, T(len(values))); err != nil {
		return err
	}

	// Write each string with its own length prefix
	for _, s := range values {
		err := PutFixedStringWithPadding(buf, s, fixedLen, padding, fromLeft)
		if err != nil {
			return nil
		}
	}
	return nil
}

func GetFixedString(buf *bytes.Buffer, fixedLen int) (string, error) {
	return GetFixedStringTrimPadding(buf, fixedLen, ' ', false)
}
func GetFixedStringTrimPadding(buf *bytes.Buffer, fixedLen int, padding rune, fromLeft bool) (string, error) {
	strBytes := make([]byte, fixedLen)
	_, err := io.ReadFull(buf, strBytes)
	if fromLeft {
		return string(bytes.TrimLeft(strBytes, string(padding))), err
	}
	return string(bytes.TrimRight(strBytes, string(padding))), err
}

func GetFixedStringList[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int) ([]string, error) {
	return GetFixedStringListTrimPadding[T](buf, fixedLen, ' ', false)
}
func GetFixedStringListTrimPadding[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int, padding rune, fromLeft bool) ([]string, error) {
	var t T
	if err := binary.Read(buf, binary.BigEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]string, 0, count)
	var err error
	for i := 0; i < count; i++ {
		str, e := GetFixedStringTrimPadding(buf, fixedLen, padding, fromLeft)
		if e != nil {
			return nil, e
		}
		result = append(result, str)
	}
	return result, err
}

func GetFixedStringListLE[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int) ([]string, error) {
	return GetFixedStringListTrimPaddingLE[T](buf, fixedLen, ' ', false)
}

func GetFixedStringListTrimPaddingLE[T constraints.Unsigned](buf *bytes.Buffer, fixedLen int, padding rune, fromLeft bool) ([]string, error) {
	var t T
	if err := binary.Read(buf, binary.LittleEndian, &t); err != nil {
		return nil, err
	}
	count := int(t)

	result := make([]string, 0, count)
	var err error
	for i := 0; i < count; i++ {
		str, e := GetFixedStringTrimPadding(buf, fixedLen, padding, fromLeft)
		if e != nil {
			return nil, e
		}
		result = append(result, str)
	}
	return result, err
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

func PutStringList[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer, values []string) error {
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

func GetStringList[T constraints.Unsigned, K constraints.Unsigned](buf *bytes.Buffer) ([]string, error) {
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
func PutObjectList[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, values []K) error {
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

func GetObjectList[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, newFn func() K) ([]K, error) {
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
func PutObjectListLE[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, values []K) error {
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

func GetObjectListLE[T constraints.Unsigned, K BinaryCodec](buf *bytes.Buffer, newFn func() K) ([]K, error) {
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
