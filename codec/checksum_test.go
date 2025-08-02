package codec_test

import (
	"bytes"
	"testing"

	"github.com/xinchentechnote/fin-proto-go/codec"
)

// helper: get service and fail if not found
func mustGet[T any](t *testing.T, algo string) any {
	svc, ok := codec.Get(algo)
	if !ok {
		t.Fatalf("expected service %s registered", algo)
	}
	return svc
}

func TestRegistryAndGet(t *testing.T) {
	// CRC16
	svcAny := mustGet[uint16](t, "CRC16")
	if svc, ok := svcAny.(*codec.Crc16ChecksumService); ok {
		sum := svc.Calc(bytes.NewBufferString("123456789"))
		if sum != 0x4B37 { // CRC16/MODBUS for "123456789"
			t.Errorf("CRC16 wrong: got %04X, want 0x4B37", sum)
		}
	} else {
		t.Fatalf("wrong type for CRC16")
	}

	// CRC32
	svcAny = mustGet[uint32](t, "CRC32")
	if svc, ok := svcAny.(*codec.Crc32ChecksumService); ok {
		sum := svc.Calc(bytes.NewBufferString("123456789"))
		if sum != 0xCBF43926 { // CRC32 IEEE
			t.Errorf("CRC32 wrong: got %08X, want CBF43926", sum)
		}
	} else {
		t.Fatalf("wrong type for CRC32")
	}

	// SSE_BIN
	svcAny = mustGet[uint32](t, "SSE_BIN")
	if svc, ok := svcAny.(*codec.SseBinChecksumService); ok {
		sum := svc.Calc(bytes.NewBuffer([]byte{1, 2, 3}))
		if sum != ((1 + 2 + 3) & 0xFF) {
			t.Errorf("SSE_BIN wrong: got %d, want %d", sum, (1+2+3)&0xFF)
		}
	} else {
		t.Fatalf("wrong type for SSE_BIN")
	}

	// SZSE_BIN
	svcAny = mustGet[int32](t, "SZSE_BIN")
	if svc, ok := svcAny.(*codec.SzseBinChecksumService); ok {
		sum := svc.Calc(bytes.NewBuffer([]byte{1, 2, 3}))
		if sum != (1+2+3)%256 {
			t.Errorf("SZSE_BIN wrong: got %d, want %d", sum, (1+2+3)%256)
		}
	} else {
		t.Fatalf("wrong type for SZSE_BIN")
	}
}

func TestRegistryDuplicate(t *testing.T) {
	ok := codec.Registry(&codec.Crc32ChecksumService{})
	if ok {
		t.Errorf("expected duplicate registry of CRC32 to fail")
	}
}

func TestRemoveAndClear(t *testing.T) {
	// Register a dummy service
	dummy := &codec.SseBinChecksumService{}
	codec.Registry(dummy)

	// Remove
	codec.Remove(dummy.Algorithm())
	if _, ok := codec.Get(dummy.Algorithm()); ok {
		t.Errorf("expected %s removed", dummy.Algorithm())
	}

	// Clear
	codec.Clear()
	if _, ok := codec.Get("CRC32"); ok {
		t.Errorf("expected cache cleared, but CRC32 still exists")
	}
}
