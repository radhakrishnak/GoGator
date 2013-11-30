package ofp10

import (

)

type OFAction struct {
	Type uint16
	Length uint16
//	Pad [4]uint8
	Port uint16 // ActionOutput only edit in future
	MaxLen uint16 // ActionOutput only edit in future
}

type ActionOutput struct {
	Type uint16
	Length uint16
	OFAction
	Port uint16
	MaxLen uint16
}

func NewOFAction() *OFAction {
	action := new(OFAction)
	action.Type = OFPAT_OUTPUT
	action.Length = 8
	action.Port = OFPP_FLOOD
	return action

}

func NewActionOutput() *ActionOutput{
	action := new(ActionOutput)
	action.Type = OFPAT_OUTPUT
	action.Length = 8
	action.Port = OFPP_FLOOD
	return action
}

const (
	OFPAT_OUTPUT = iota
	OFPAT_SET_VLAN_VID
	OFPAT_SET_VLAN_PCP
	OFPAT_STRIP_VLAN
	OFPAT_SET_DL_SRC
	OFPAT_SET_DL_DST
	OFPAT_SET_NW_SRC
	OFPAT_SET_NW_DST
	OFPAT_SET_NW_TOS
	OFPAT_SET_TP_SRC
	OFPAT_SET_TP_DST
	OFPAT_ENQUEUE
	OFPAT_VENDOR = 0xffff
)