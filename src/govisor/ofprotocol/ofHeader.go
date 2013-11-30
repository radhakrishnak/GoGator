package ofp10

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
)

// OpenFlow Specification 1.0.0   Version=0x01
const (
	/* Immutable messages. */
	T_HELLO        = iota // Symmetric message  //0
	T_ERROR               // Symmetric message	//1
	T_ECHO_REQUEST        // Symmetric message
	T_ECHO_REPLY          // Symmetric message
	T_VENDOR              // Symmetric message

	/* Switch configuration messages. */
	T_FEATURES_REQUEST   // Controller/Switch message
	T_FEATURES_REPLY     // Controller/Switch message
	T_GET_CONFIG_REQUEST // Controller/Switch message
	T_GET_CONFIG_REPLY   // Controller/Switch message
	T_SET_CONFIG         // Controller/Switch message

	/* Asynchronous messages. */
	T_PACKET_IN    /* Asynchronous messages. */
	T_FLOW_REMOVED /* Asynchronous messages. */
	T_PORT_STATUS  /* Asynchronous messages. */

	/* Controller command messages. */
	T_PACKET_OUT // Controller/Switch message
	T_FLOW_MOD   // Controller/Switch message
	T_PORT_MOD   // Controller/Switch message

	/* Statistics messages. */
	T_STATS_REQUEST // Controller/Switch message
	T_STATS_REPLY   // Controller/Switch message

	/* Barrier messages. */
	T_BARRIER_REQUEST // Controller/Switch message
	T_BARRIER_REPLY   // Controller/Switch message

	/* Queue Configuration messages. */
	T_QUEUE_GET_CONFIG_REQUEST // Controller/Switch message
	T_QUEUE_GET_CONFIG_REPLY   // Controller/Switch message
)

type Header struct {
	Version uint8
	Type    uint8
	Length  uint16
	XID     uint32
}

func (h *Header) GetHeader() *Header {
	return h
}

func (h *Header) Read(b_Header []byte) (n int, err error) {
	bufHeader := new(bytes.Buffer)
	binary.Write(bufHeader, binary.BigEndian, h) ////write h into bufHeader
	n, err = bufHeader.Read(b_Header)
	if err != nil {
		log.Println("ERROR:SendEchoReply::ofHeader.Read(): ", err)
		return
	}
	return n, io.EOF
}

func (h *Header) Write(b []byte) (n int, err error) {
	buf := bytes.NewBuffer(b)
	binary.Read(buf, binary.BigEndian, h)
	return 8, err
}

var NewHeader func() *Header = newHeaderGenerator()
func newHeaderGenerator() func() *Header {
	var xid uint32 = 1
	return func() *Header {
		p := new(Header)
		p.Version = 1
		p.Type = 0
		p.Length = 8
		p.XID = xid
		xid += 1
		return p
	}
}

//Basic of protocols based only Header
func Hello() *Header {
	h := NewHeader()
	h.Type = T_HELLO
	h.Version = 1
	h.Length = 8
	h.XID = 1
	return h
}

func ErrorHeader() *Header {
	h := NewHeader()
	h.Type = T_ERROR
	return h
}

func EchoRequest() *Header {
	h := NewHeader()
	h.Type = T_ECHO_REQUEST
	return h
}
func EchoReply() *Header {
	h := NewHeader()
	h.Type = T_ECHO_REPLY
	return h
}
func Vendor() *Header {
	h := NewHeader()
	h.Type = T_VENDOR
	return h
}

// FeaturesRequest constructor
func FeaturesRequest() *Header {
	h := NewHeader()
	h.Type = T_FEATURES_REQUEST
	return h
}

func FeaturesReply() *Header {
	h := NewHeader()
	h.Type = T_FEATURES_REPLY
	return h
}

func GetConfigRequest() *Header {
	h := NewHeader()
	h.Type = T_GET_CONFIG_REQUEST
	return h
}
func GetConfigReply() *Header {
	h := NewHeader()
	h.Type = T_GET_CONFIG_REPLY
	return h
}
func SetConfig() *Header {
	h := NewHeader()
	h.Type = T_SET_CONFIG
	return h
}

func FlowRemoved() *Header {
	h := NewHeader()
	h.Type = T_FLOW_REMOVED
	return h
}
func PortMod() *Header {
	h := NewHeader()
	h.Type = T_PORT_MOD
	return h
}

func BarrierRequest() *Header {
	h := NewHeader()
	h.Type = T_BARRIER_REQUEST
	return h
}
func BarrierReply() *Header {
	h := NewHeader()
	h.Type = T_BARRIER_REPLY
	return h
}
func QueueGetConfigRequest() *Header {
	h := NewHeader()
	h.Type = T_QUEUE_GET_CONFIG_REQUEST
	return h
}
func QueueGetConfigReply() *Header {
	h := NewHeader()
	h.Type = T_QUEUE_GET_CONFIG_REPLY
	return h
}

// Msg
type IPacket interface {
	io.ReadWriter
	GetHeader() *Header
}

type MsgPacket struct {
	Data IPacket
	DPID string
}
type MsgByte struct {
	Data []byte
	DPID string
}

