package codec_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xinchentechnote/fin-proto-go/codec"
)

func TestWriteAndReadStringLE(t *testing.T) {
	var buf bytes.Buffer
	original := "hello world"

	codec.WriteStringLE[uint16](&buf, original)

	decoded, err := codec.ReadStringLE[uint16](&buf)

	require.NoError(t, err)
	require.Equal(t, original, decoded)
}

func TestWriteAndReadFixedString(t *testing.T) {
	var buf bytes.Buffer
	original := "abc"
	codec.WriteFixedString(&buf, original, 8)

	decoded, err := codec.ReadFixedString(&buf, 8)

	require.NoError(t, err)
	require.Equal(t, original, decoded)
}

func TestFixedStringOverflow(t *testing.T) {
	var buf bytes.Buffer
	original := "toolongstring"
	codec.WriteFixedString(&buf, original, 5)

	decoded, err := codec.ReadFixedString(&buf, 5)

	require.NoError(t, err)
	require.Equal(t, original[:5], decoded)
}

func TestWriteAndReadStringListLE(t *testing.T) {
	original := []string{"foo", "bar", "baz"}

	var buf bytes.Buffer
	codec.WriteStringListLE[uint16, uint16](&buf, original)
	decoded, err := codec.ReadStringListLE[uint16, uint16](&buf)
	require.NoError(t, err)
	require.Equal(t, original, decoded)
}

func TestReadBasicType(t *testing.T) {
	order := binary.BigEndian

	t.Run("uint8", func(t *testing.T) {
		var original uint8 = 123
		buf := new(bytes.Buffer)
		err := binary.Write(buf, order, original)
		if err != nil {
			t.Fatalf("failed to write: %v", err)
		}
		val, err := codec.ReadBasicType[uint8](buf)
		if err != nil || val != original {
			t.Errorf("expected %v, got %v (err=%v)", original, val, err)
		}
	})

	t.Run("int16", func(t *testing.T) {
		var original int16 = -4567
		buf := new(bytes.Buffer)
		err := binary.Write(buf, order, original)
		if err != nil {
			t.Fatalf("failed to write: %v", err)
		}
		val, err := codec.ReadBasicType[int16](buf)
		if err != nil || val != original {
			t.Errorf("expected %v, got %v (err=%v)", original, val, err)
		}
	})

	t.Run("uint32", func(t *testing.T) {
		var original uint32 = 987654321
		buf := new(bytes.Buffer)
		err := binary.Write(buf, order, original)
		if err != nil {
			t.Fatalf("failed to write: %v", err)
		}
		val, err := codec.ReadBasicType[uint32](buf)
		if err != nil || val != original {
			t.Errorf("expected %v, got %v (err=%v)", original, val, err)
		}
	})

	t.Run("float64", func(t *testing.T) {
		var original float64 = 3.1415926535
		buf := new(bytes.Buffer)
		err := binary.Write(buf, order, original)
		if err != nil {
			t.Fatalf("failed to write: %v", err)
		}
		val, err := codec.ReadBasicType[float64](buf)
		if err != nil || val != original {
			t.Errorf("expected %v, got %v (err=%v)", original, val, err)
		}
	})
}
