package ofp10

import (
	"bytes"
	"encoding/binary"
	"io"
)

type PacketOut struct {
	Header       *Header
	BufferID     uint32
	InPort       uint16
	ActionLength uint16
	Action       *OFAction
	PacketData   []byte
}

func NewPacketOut() *PacketOut {
	p := new(PacketOut)
	p.Header = NewHeader()
	p.Header.Type = T_PACKET_OUT
	p.Header.Length = 16		//default when no action
	return p
}

func (p *PacketOut) SetInport(InPort uint16) {
	p.InPort = InPort
}

func (p *PacketOut) SetBufferID(BufferID uint32) {
	p.BufferID = BufferID

}

func (p *PacketOut) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(p.Header)
	binary.Write(buf, binary.BigEndian, p.BufferID)
	binary.Write(buf, binary.BigEndian, p.InPort)
	binary.Write(buf, binary.BigEndian, p.ActionLength)
	binary.Write(buf, binary.BigEndian, p.Action)
	binary.Write(buf, binary.BigEndian, p.PacketData)
	n, err = buf.Read(b)
	if n == 0 {
		return
	}
	return n, io.EOF
}
