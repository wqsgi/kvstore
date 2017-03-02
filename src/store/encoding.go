package store

import (
	"encoding/binary"
	"github.com/klauspost/crc32"
)

const (
	HFileHaderSize = 20
)

const(
	Magic uint32= 0xED0CDAEE
)
// crc32(4) timestamp 8 keysize 4 valuesize 4
func encodingHFileData(key, value []byte, timestamp uint64, version uint32) []byte {
	keySize := uint16(len(key))
	valueSize := uint16(len(value))

	buf := make([]byte, HFileHaderSize+keySize+valueSize)

	//
	binary.LittleEndian.PutUint64(buf[4:12], timestamp)
	binary.LittleEndian.PutUint16(buf[12:14], keySize)
	binary.LittleEndian.PutUint16(buf[14:16], valueSize)
	binary.LittleEndian.PutUint32(buf[16:20], version)

	copy(buf[HFileHaderSize:HFileHaderSize+keySize], key)
	copy(buf[HFileHaderSize+keySize:HFileHaderSize+keySize+valueSize], value)

	binary.LittleEndian.PutUint32(buf[0:4], crc32.ChecksumIEEE(buf[4:]))
	return buf
}

func decodingHFileData(value []byte)*HLog{



}