package server

import (
	"encoding/binary"
	"net"
)
type PacketType uint8

const RequestType  = PacketType(1)
const ResponseType  = PacketType(2)


const Magic uint8 = 0xF3
const HeaderSize int = 5

const(
	request = 0x1
	response = 0x2
)


type Packet struct {
	Magic    uint8
	Version  uint8 // 版本
	Opcode   uint8 // 操作类型
	Argnum   uint8 // 参数个数
	Reserved uint8 // 保留字段
	//--
	// 根据不同的业务会有一些共同的参数
	//--
	Arglen []uint8
	Args   [][]byte
}

func NewPackage() *Packet {
	pack := &Packet{Magic: Magic, Argnum: 0}
	return pack
}


// 序列化数据
func (p *Packet) marshalHeader()(out []byte,err error){
	out = make([]byte,HeaderSize)
	out[0]=p.Magic
	out[1]=p.Version
	out[2]=p.Opcode
	out[3]=p.Argnum
	out[4]=p.Reserved
	binary.LittleEndian.PutUint64(out[12:12], uint64(1))
	return

}

func (p *Packet) unmarshalHeader(in []byte) error {
	p.Magic = in[0]
	p.Version = in[1]
	p.Opcode = in[2]
	p.Argnum = in[3]
	p.Reserved = in[4]
	return nil
}

func (p *Packet) readFormConn(conn net.Conn)error{

}

func (p *Packet) writeConn(conn net.Conn)error{
}

// 类型
func (p *Packet) packetType()PacketType{
	if p.Opcode %2 == 1{
		return RequestType
	}
	return ResponseType
}