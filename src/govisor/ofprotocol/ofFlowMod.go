package ofprotocol

import (
	"bytes"
	"encoding/binary"
	"io"
)

type OFFlowMod struct {
	Header  *Header
	OFMatch *OFMatch
	Cookie  uint64

	Command     uint16
	IdleTimeout uint16
	HardTimeout uint16
	Priority    uint16
	BufferID    uint32
	OutPort     uint16
	Flags       uint16
	Action      *OFAction
}

func NewFlowMod() *OFFlowMod {
	f := new(OFFlowMod)
	f.Header = NewHeader()
	f.Header.Type = T_FLOW_MOD
	f.OFMatch = new(OFMatch)
	// Add a generator for f.Cookie here
	f.Cookie = 0

	f.Command = OFPFC_ADD
	f.IdleTimeout = 0
	f.HardTimeout = 0
	// Add a priority gen here
	f.Priority = 1000
	f.BufferID = 0xffffffff
	f.OutPort = OFPP_NONE
	f.Flags = 0
	//f.Actions = make([]Packetish, 0)
	return f
}

//func (fm *OFFlowMod) SetInport (InPort uint16)  {
//	fm.InPort = InPort
//
//}

//func (fm *OFFlowMod) SetBufferID (BufferID uint32)  {
//	fm.BufferID = BufferID
//
//}
//func (fm *OFFlowMod) SetCommand (Command uint16)  {
//	fm.Command = Command
//
//}
func (f *OFFlowMod) ComputeLength() (n uint16) {
	//	for _, v := range f.Actions {
	//		n += v.Len()
	//	}
	n = 72 + 8 // 72 + actionlength
	return n
}

func (f *OFFlowMod) Read(b []byte) (n int, err error) {

	f.Header.Length = f.ComputeLength()
	buf := new(bytes.Buffer)
	buf.ReadFrom(f.Header)
	//	fmt.Println("FlowMod Read1: Header: ", buf.Bytes())
	buf.ReadFrom(f.OFMatch)
	//	fmt.Println("FlowMod Read2: Match: ", buf.Bytes())
	binary.Write(buf, binary.BigEndian, f.Cookie)
	//	fmt.Println("FlowMod Read3: Cookie: ", buf.Bytes())
	binary.Write(buf, binary.BigEndian, f.Command)
	binary.Write(buf, binary.BigEndian, f.IdleTimeout)
	binary.Write(buf, binary.BigEndian, f.HardTimeout)
	binary.Write(buf, binary.BigEndian, f.Priority)
	binary.Write(buf, binary.BigEndian, f.BufferID)
	binary.Write(buf, binary.BigEndian, f.OutPort)
	binary.Write(buf, binary.BigEndian, f.Flags)
	binary.Write(buf, binary.BigEndian, f.Action.Type)
	binary.Write(buf, binary.BigEndian, f.Action.Length)
	binary.Write(buf, binary.BigEndian, f.Action.Port)
	binary.Write(buf, binary.BigEndian, f.Action.MaxLen)

	//	fmt.Println("FlowMod Read4: all: ", buf.Bytes())

	n, err = buf.Read(b)
	if err != nil {
		return
	}
	return n, io.EOF
}

// ofp_flow_mod_command 1.0
const (
	OFPFC_ADD = iota // OFPFC_ADD == 0
	OFPFC_MODIFY
	OFPFC_MODIFY_STRICT
	OFPFC_DELETE
	OFPFC_DELETE_STRICT
)

// ofp_flow_mod_flags 1.0
const (
	OFPFF_SEND_FLOW_REM = 1 << 0
	OFPFF_CHECK_OVERLAP = 1 << 1
	OFPFF_EMERG         = 1 << 2
)
