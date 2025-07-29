package codec

import (
	"bytes"
	"hash/crc32"
	"sync"
)

type ChecksumService[T any, R any] interface {
	Algorithm() string
	Calc(input T) R
}

type ChecksumServiceContext struct {
	mu    sync.RWMutex
	cache map[string]any
}

var checksumServiceContext = &ChecksumServiceContext{
	cache: make(map[string]any),
}

func Registry(service any) bool {
	if cs, ok := service.(interface{ Algorithm() string }); ok {
		checksumServiceContext.mu.Lock()
		defer checksumServiceContext.mu.Unlock()
		if _, exists := checksumServiceContext.cache[cs.Algorithm()]; exists {
			return false
		}
		checksumServiceContext.cache[cs.Algorithm()] = service
		return true
	}
	return false
}

func Get(algorithm string) (any, bool) {
	checksumServiceContext.mu.RLock()
	defer checksumServiceContext.mu.RUnlock()
	if service, exists := checksumServiceContext.cache[algorithm]; exists {
		return service, exists
	}
	return nil, false
}

func Remove(algorithm string) {
	checksumServiceContext.mu.Lock()
	defer checksumServiceContext.mu.Unlock()
	delete(checksumServiceContext.cache, algorithm)
}

func Clear() {
	checksumServiceContext.mu.Lock()
	defer checksumServiceContext.mu.Unlock()
	checksumServiceContext.cache = make(map[string]any)
}

type Crc16ChecksumService struct{}

// c Crc16ChecksumService
func (c *Crc16ChecksumService) Algorithm() string {
	return "CRC16"
}

func (c *Crc16ChecksumService) Calc(data *bytes.Buffer) uint16 {
	var crc uint16 = 0xFFFF
	for _, b := range data.Bytes() {
		crc ^= uint16(b)
		for i := 0; i < 8; i++ {
			if crc&0x0001 != 0 {
				crc = (crc >> 1) ^ 0xA001
			} else {
				crc >>= 1
			}
		}
	}
	return crc
}

type Crc32ChecksumService struct{}

func (c *Crc32ChecksumService) Algorithm() string {
	return "CRC32"
}
func (c *Crc32ChecksumService) Calc(data *bytes.Buffer) uint32 {
	return crc32.ChecksumIEEE(data.Bytes())
}

type SseBinChecksumService struct{}

func (c *SseBinChecksumService) Algorithm() string {
	return "SSE_BIN"
}
func (c *SseBinChecksumService) Calc(data *bytes.Buffer) uint32 {
	var checksum uint32
	for _, b := range data.Bytes() {
		checksum = (checksum + uint32(b)) & 0xFF
	}
	return checksum
}

type SzseBinChecksumService struct{}

func (c *SzseBinChecksumService) Algorithm() string {
	return "SZSE_BIN"
}
func (c *SzseBinChecksumService) Calc(data *bytes.Buffer) uint32 {
	var checksum uint32
	for _, b := range data.Bytes() {
		checksum += uint32(b)
	}
	return checksum % 256
}

func init() {
	Registry(&Crc16ChecksumService{})
	Registry(&Crc32ChecksumService{})
	Registry(&SseBinChecksumService{})
	Registry(&SzseBinChecksumService{})
}
