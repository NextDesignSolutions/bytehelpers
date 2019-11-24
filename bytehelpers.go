package bytehelpers

import (
	"bytes"
	"encoding/binary"
)

func BytesToUint32s(data []byte, order binary.ByteOrder) ([]uint32, error) {
	num_uint32 := len(data) / 4
	u32 := make([]uint32, num_uint32)
	rbuf := bytes.NewReader(data)
	err := binary.Read(rbuf, order, u32)
	if err != nil {
		return nil, err
	}
	return u32, err
}

func BytesToUint64s(data []byte, order binary.ByteOrder) ([]uint64, error) {
	num_uint64 := len(data) / 8
	u64 := make([]uint64, num_uint64)
	rbuf := bytes.NewReader(data)
	err := binary.Read(rbuf, order, u64)
	if err != nil {
		return nil, err
	}
	return u64, err
}

func Uint32sToBytes(data []uint32, order binary.ByteOrder) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func Uint64sToBytes(data []uint64, order binary.ByteOrder) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, order, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
