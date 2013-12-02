package ofp10

import (
	"bytes"
	"encoding/binary"
	//"fmt"
	"io"
)

var MINIMUM_LENGTH uint16 = 48
var OFP_ETH_ALEN int = 6
var OFP_MAX_PORT_NAME_LEN int = 16

type OFPortConfig struct {
	Value int
}

func (pc *OFPortConfig) NewOFPortConfig(value int) {
	pc.Value = value
}

func (pc *OFPortConfig) getValue() int {
	return pc.Value
}

type OFPortState struct {
	Value int
}

func (ps *OFPortState) NewOFPortState(value int) {
	ps.Value = value
}

func (ps *OFPortState) getValue() int {
	return ps.Value
}

type OFPortFeatures struct {
	Value int
}

func (pf *OFPortFeatures) NewOFPortFeatures(value int) {
	pf.Value = value
}

func (pf *OFPortFeatures) getValue() int {
	return pf.Value
}

type OFPhysicalPort struct {
	//PortNo     uint32
	PortNo     uint16
	HWAddr     []byte
	Name       string
	Config     uint32
	State      uint32
	Curr       uint32
	Advertised uint32
	Supported  uint32
	Peer       uint32
	//CurrSpeed  uint32
	//MaxSpeed   uint32
}

func (p *OFPhysicalPort) Read(b []byte) (n int, err error) {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, p)
	n, err = buf.Read(b)
	if err != nil {
		return
	}
	return n, io.EOF
}

func (p *OFPhysicalPort) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	//fmt.Println("***buf*** ", buf.Bytes()[0:])
	//fmt.Println("Getting sw features: ", d_buf.Bytes()[0:32])
	err = binary.Read(buf, binary.BigEndian, &p.PortNo)
	if err != nil {
		return
	}

	//fmt.Println("***PortNo*** ", p.PortNo)
	n += 2
	//buf.Next(4) //padding
	p.HWAddr = make([]byte, OFP_ETH_ALEN)
	err = binary.Read(buf, binary.BigEndian, &p.HWAddr)
	if err != nil {
		return
	}
	//fmt.Println("***HWAddr*** ", p.HWAddr)
	n += 6
	//buf.Next(2) //padding
	name := make([]byte, OFP_MAX_PORT_NAME_LEN)
	err = binary.Read(buf, binary.BigEndian, &name)
	p.Name = string(name)
	//
	//err = binary.Read(buf, binary.BigEndian, &p.Name)
	if err != nil {
		return
	}
	//fmt.Println("***Name*** ", p.Name)
	n += 16
	err = binary.Read(buf, binary.BigEndian, &p.Config)
	if err != nil {
		return
	}
	//fmt.Println("***Config*** ", p.Config)
	n += 4
	err = binary.Read(buf, binary.BigEndian, &p.State)
	if err != nil {
		return
	}
	//fmt.Println("***State*** ", p.State)
	n += 4
	err = binary.Read(buf, binary.BigEndian, &p.Curr)
	if err != nil {
		return
	}
	//fmt.Println("***Curr*** ", p.Curr)
	n += 4
	err = binary.Read(buf, binary.BigEndian, &p.Advertised)
	if err != nil {
		return
	}
	//fmt.Println("***Advertised*** ", p.Advertised)
	n += 4
	err = binary.Read(buf, binary.BigEndian, &p.Supported)
	if err != nil {
		return
	}
	//fmt.Println("***Supported*** ", p.Supported)
	n += 4
	err = binary.Read(buf, binary.BigEndian, &p.Peer)
	if err != nil {
		return
	}
	//fmt.Println("***Peer*** ", p.Peer)
	n += 4
	//err = binary.Read(buf, binary.BigEndian, &p.CurrSpeed)
	//if err != nil {
	//	return
	//}
	//fmt.Println("***CurrSpeed*** ", p.CurrSpeed)
	//n += 4
	//err = binary.Read(buf, binary.BigEndian, &p.MaxSpeed)
	//if err != nil {
	//	return
	//}
	//fmt.Println("***MaxSpeed*** ", p.MaxSpeed)
	//n += 4
	return n, err
}

const (
	OFPPC_PORT_DOWN    = 1 << 0
	OFPPC_NO_STP       = 1 << 1
	OFPPC_NO_RECV      = 1 << 2
	OFPPC_NO_RECV_STP  = 1 << 3
	OFPPC_NO_FLOOD     = 1 << 4
	OFPPC_NO_FWD       = 1 << 5
	OFPPC_NO_PACKET_IN = 1 << 6

	OFPPS_LINK_DOWN   = 1 << 0
	OFPPS_STP_LISTEN  = 0 << 8
	OFPPS_STP_LEARN   = 1 << 8
	OFPPS_STP_FORWARD = 2 << 8
	OFPPS_STP_BLOCK   = 3 << 8
	OFPPS_STP_MASK    = 3 << 8

	OFPPF_10MB_HD    = 1 << 0
	OFPPF_10MB_FD    = 1 << 1
	OFPPF_100MB_HD   = 1 << 2
	OFPPF_100MB_FD   = 1 << 3
	OFPPF_1GB_HD     = 1 << 4
	OFPPF_1GB_FD     = 1 << 5
	OFPPF_10GB_FD    = 1 << 6
	OFPPF_COPPER     = 1 << 7
	OFPPF_FIBER      = 1 << 8
	OFPPF_AUTONEG    = 1 << 9
	OFPPF_PAUSE      = 1 << 10
	OFPPF_PAUSE_ASYM = 1 << 11
)
