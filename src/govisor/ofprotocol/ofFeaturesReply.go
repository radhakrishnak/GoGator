package ofprotocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"govisor/utils"
)

const (
	OFPC_FLOW_STATS   = 1 << 0
	OFPC_TABLE_STATS  = 1 << 1
	OFPC_PORT_STATS   = 1 << 2
	OFPC_STP          = 1 << 3
	OFPC_RESERVED     = 1 << 4
	OFPC_IP_REASM     = 1 << 5
	OFPC_QUEUE_STATS  = 1 << 6
	OFPC_ARP_MATCH_IP = 1 << 7
)

type OFCapabilities struct {
	Value int
}

func (c *OFCapabilities) GetValue() int {
	return c.Value
}

func (c *OFCapabilities) SetValue(v int) {
	c.Value = v
}

var MINIMUM_LENGTH_FR uint16 = 32 //FR: Features Reply

type OFFeaturesReply struct {
	Header Header
	//DataPathId   uint64
	DataPathId   []uint8
	Buffers      uint32
	Tables       uint8
	Capabilities uint32
	Actions      uint32
	Ports        []OFPhysicalPort
}

func NewFeaturesReply() *OFFeaturesReply {
	f := new(OFFeaturesReply)
	//f.Header = NewHeader()
	f.Header.Type = T_FEATURES_REPLY
	f.Header.Length = MINIMUM_LENGTH_FR
	return f
}

func (f *OFFeaturesReply) ComputeLength(n int16) int16 {
	length := util.Uint16_Int16(MINIMUM_LENGTH_FR)
	return length
}

func (f *OFFeaturesReply) GetHeader() *Header {
	return &f.Header
}

func (f *OFFeaturesReply) WriteTo(b []byte) (n int, err error) { //writeTo

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, f) //write f into buf
	n, err = buf.Read(b)                   //read buf
	if n == 0 {
		return
	}
	return n, io.EOF
}

func (f *OFFeaturesReply) Read(b []byte) (n int, err error) { //equal to writeTo

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, f) //write f into buf
	n, err = buf.Read(b)                   //read buf
	if n == 0 {
		return
	}
	return n, io.EOF
}

func (h *Header) ReadFrom(b []byte) (n int, err error) { //readFrom

	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, h) //read buf into h
	return 8, err
}

func (f *OFFeaturesReply) ReadFrom(b []byte) (n int, err error) { //readFrom

	buf := bytes.NewBuffer(b)
	n, err = f.Header.Write(buf.Next(8))
	if n == 0 {
		return
	}
	if err = binary.Read(buf, binary.BigEndian, &f.DataPathId); err != nil {
		return
	}
	n += 8
	if err = binary.Read(buf, binary.BigEndian, &f.Buffers); err != nil {
		return
	}
	n += 4
	if err = binary.Read(buf, binary.BigEndian, &f.Tables); err != nil {
		return
	}
	n += 1
	n += 3 //padding
	if err = binary.Read(buf, binary.BigEndian, &f.Capabilities); err != nil {
		return
	}
	n += 4
	if err = binary.Read(buf, binary.BigEndian, &f.Actions); err != nil {
		return
	}
	n += 4

	//Ports
	portCount := (f.Header.Length - 32) / MINIMUM_LENGTH //n is 32
	f.Ports = make([]OFPhysicalPort, 0)

	for i := 0; i < int(portCount); i++ {
		m := 0
		port := OFPhysicalPort{}
		m, err := port.Write(b[n:])
		if m == 0 {
			return m, err
		}
		f.Ports = append(f.Ports, port)
		n += m
	}

	return
}

//func (h *Header) Write(b []byte) (n int, err error) { //readFrom

//	buf := bytes.NewBuffer(b)
//	binary.Read(buf, binary.BigEndian, h) //read buf into h
//	return 8, err
//}

func (f *OFFeaturesReply) Write(b []byte) (n int, err error) { //readFrom

	buf := bytes.NewBuffer(b)
	n, err = f.Header.Write(buf.Next(8))
	if n == 0 {
		return
	}
	f.DataPathId = make([]uint8, 8)
	if err = binary.Read(buf, binary.BigEndian, &f.DataPathId); err != nil {
		return
	}
	n += 8

	if err = binary.Read(buf, binary.BigEndian, &f.Buffers); err != nil {
		return
	}
	n += 4

	if err = binary.Read(buf, binary.BigEndian, &f.Tables); err != nil {
		return
	}
	n += 1
	n += 3 //padding
	buf.Next(3)
	if err = binary.Read(buf, binary.BigEndian, &f.Capabilities); err != nil {
		return
	}
	n += 4
	if err = binary.Read(buf, binary.BigEndian, &f.Actions); err != nil {
		return
	}
	n += 4

	portCount := (f.Header.Length - 32) / MINIMUM_LENGTH //n is 32
	f.Ports = make([]OFPhysicalPort, 0)

	for i := 0; i < int(portCount); i++ {
		m := 0
		port := OFPhysicalPort{}
		m, err := port.Write(b[n:])
		fmt.Println("OFFeaturesReply: ", port)
			
		if m == 0 {
			return m, err
		}
		f.Ports = append(f.Ports, port)
		n += m
	}
	return
}
