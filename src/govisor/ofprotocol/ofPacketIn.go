package ofp10

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	//"golfer/util"
	"govisor/pkt"
)

type PacketIn struct {
	Header      Header
	BufferID    uint32
	TotalLength uint16
	InPort      uint16
	Reason      uint8
	Data        packet.Ethernet
}

func NewPacketIn() *PacketIn {
	p := new(PacketIn)
	p.Header.Type = T_PACKET_IN
	return p

}

func (p PacketIn) GetHeader() *Header {
	return &(p.Header)
}

func (p PacketIn) Read(b []byte) (n int, err error) {

	b[0]=1
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p)
	//binary.Write(buf[0:1], binary.BigEndian, p.InPort+1)
	n, err = buf.Read(b)
	if n == 0 {
		return
	}
	return n, io.EOF
}

func (p *PacketIn) Write(b []byte) (n int, err error) {
	fmt.Println("PacketIn bytes: ", b)
	buf := bytes.NewBuffer(b)
	n, err = p.Header.Write(buf.Next(8))
	p.Header.Type = T_PACKET_IN
	if n == 0 {
		return
	}
	if err = binary.Read(buf, binary.BigEndian, &p.BufferID); err != nil {
		return
	}
	p.BufferID = 3
	fmt.Println("PacketIn p.BufferID: ", p.BufferID)
	n += 4
	if err = binary.Read(buf, binary.BigEndian, &p.TotalLength); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &p.InPort); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &p.Reason); err != nil {
		return
	}
	n += 1
	n += 1 //pad
	//TODO::Parse Data
	//p.Data = util.Ethernet{}
	//fmt.Println("PacketIn buffer: ", b[n:])
	m, err := p.Data.Write(b[n:])	//ethernet
	if m == 0 {
		fmt.Println("PacketIn p.Data.Write(): ", err.Error())
		return m, err
	}
	//fmt.Println("PacketIn p.Data: ", p.Data)
	//fmt.Println("PacketIn m: ", m)
	n += m
	return n, err
}

const (
	OFPR_NO_MATCH = iota
	OFPR_ACTION
)
