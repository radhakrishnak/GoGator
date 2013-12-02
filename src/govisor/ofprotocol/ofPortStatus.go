package ofp10

import (
	"bytes"
	"encoding/binary"
	"io"
	//"networkUtils"
)

type PortStatus struct {
	Header Header
	Reason uint8
	Pad    [7]uint8
	Desc   OFPhysicalPort
}

func NewPortStatus() *PortStatus {
	p := new(PortStatus)
	p.Header.Type = T_PORT_STATUS
	return p

}

func (p *PortStatus) GetHeader() *Header {
	return &p.Header
}

func (p *PortStatus) Read(b []byte) (n int, err error) {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p)
	n, err = buf.Read(b)
	if n == 0 {
		return
	}
	return n, io.EOF
}

func (p *PortStatus) Write(b []byte) (n int, err error) {

	buf := bytes.NewBuffer(b)
	n, err = p.Header.Write(buf.Next(8))
	if n == 0 {
		return
	}
	if err = binary.Read(buf, binary.BigEndian, &p.Reason); err != nil {
		return
	}
	n += 1
	//Read pading

	n += 7
	//
	m := 0
	//p.desc = ofPhyPort
	if m, err := p.Desc.Write(b[n:]); m == 0 {
		return m, err
	}
	n += m

		/*if err = binary.Read(buf, binary.BigEndian, &p.TotalLength); err != nil {
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
		//TODO::Parse Data
		m := 0
		p.Data = networkUtils.Ethernet{}
		if m, err := p.Data.Write(b[n:]); m == 0 {
			return m, err
		}
		n += m*/
	return
}

const (
	OFPPR_ADD = iota
	OFPPR_DELETE
	OFPPR_MODIFY
)
