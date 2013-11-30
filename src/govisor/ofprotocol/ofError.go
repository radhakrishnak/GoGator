package ofp10

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Error struct {
	Header Header
	Type   uint16
	Code   uint16
	Data   uint8
}

func NewError() *Error {
	e := new(Error)
	e.Header.Type = T_ERROR
	return e

}

func (e *Error) GetHeader() *Header {
	return &e.Header
}

func (e *Error) Read(b []byte) (n int, err error) {

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, e)
	n, err = buf.Read(b)
	if n == 0 {
		return
	}
	return n, io.EOF
}

//func (h *Header) Write(b []byte) (n int, err error) {

//	buf := bytes.NewBuffer(b)
//	binary.Read(buf, binary.BigEndian, h)
//	return 8, err
//}

func (e *Error) Write(b []byte) (n int, err error) {

	buf := bytes.NewBuffer(b)
	n, err = e.Header.Write(buf.Next(8))
	if n == 0 {
		return
	}
	if err = binary.Read(buf, binary.BigEndian, &e.Type); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &e.Code); err != nil {
		return
	}
	n += 2
	if err = binary.Read(buf, binary.BigEndian, &e.Data); err != nil {
		return
	}
	n += 1
	////TODO::Parse Data
	//m := 0
	//if m, err := e.Data.Write(b[n:]); m == 0 {
	//	return m, err
	//}
	//n += m
	return
}

const (
	OFPET_HELLO_FAILED         = iota /* Hello protocol failed. */
	OFPET_BAD_REQUEST                 /* Request was not understood. */
	OFPET_BAD_ACTION                  /* Error in action description. */
	OFPET_BAD_INSTRUCTION             /* Error in instruction list. */
	OFPET_BAD_MATCH                   /* Error in match. */
	OFPET_FLOW_MOD_FAILED             /* Problem modifying flow entry. */
	OFPET_GROUP_MOD_FAILED            /* Problem modifying group entry. */
	OFPET_PORT_MOD_FAILED             /* Port mod request failed. */
	OFPET_TABLE_MOD_FAILED            /* Table mod request failed. */
	OFPET_QUEUE_OP_FAILED             /* Queue operation failed. */
	OFPET_SWITCH_CONFIG_FAILED        /* Switch config request failed. */
)
