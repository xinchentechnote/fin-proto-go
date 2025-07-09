package codec_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xinchentechnote/fin-proto-go/internal/codec"
)

func TestPutAndGetStringLE(t *testing.T) {
	var buf bytes.Buffer
	original := "hello world"

	codec.PutStringLE[uint16](&buf, original)

	decoded, err := codec.GetStringLE[uint16](&buf)

	require.NoError(t, err)
	require.Equal(t, original, decoded)
}

func TestPutAndGetFixedString(t *testing.T) {
	var buf bytes.Buffer
	original := "abc"
	codec.PutFixedString(&buf, original, 8)

	decoded, err := codec.GetFixedString(&buf, 8)

	require.NoError(t, err)
	require.Equal(t, original, decoded)
}

func TestFixedStringOverflow(t *testing.T) {
	var buf bytes.Buffer
	original := "toolongstring"
	codec.PutFixedString(&buf, original, 5)

	decoded, err := codec.GetFixedString(&buf, 5)

	require.NoError(t, err)
	require.Equal(t, original[:5], decoded)
}

func TestPutAndGetStringListLE(t *testing.T) {
	original := []string{"foo", "bar", "baz"}

	var buf bytes.Buffer
	codec.PutStringListLE[uint16, uint16](&buf, original)
	decoded, err := codec.GetStringListLE[uint16, uint16](&buf)
	require.NoError(t, err)
	require.Equal(t, original, decoded)
}

func TestGetBasicType(t *testing.T) {
	order := binary.LittleEndian

	t.Run("uint8", func(t *testing.T) {
		var original uint8 = 123
		buf := new(bytes.Buffer)
		err := binary.Write(buf, order, original)
		if err != nil {
			t.Fatalf("failed to write: %v", err)
		}
		val, err := codec.GetBasicType[uint8](buf, order)
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
		val, err := codec.GetBasicType[int16](buf, order)
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
		val, err := codec.GetBasicType[uint32](buf, order)
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
		val, err := codec.GetBasicType[float64](buf, order)
		if err != nil || val != original {
			t.Errorf("expected %v, got %v (err=%v)", original, val, err)
		}
	})
}
