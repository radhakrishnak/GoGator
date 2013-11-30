package ofp10

import ()

type OFPort struct {
	Value int
}

func (p *OFPort) NewOFPortConfig(value int) {
	p.Value = value
}

func (p *OFPort) getValue() int {
	return p.Value
}

const (
	OFPP_MAX = 0Xff00

	OFPP_IN_PORT = 0xfff8
	OFPP_TABLE   = 0xfff9

	OFPP_NORMAL = 0xfffa
	OFPP_FLOOD  = 0xfffb

	OFPP_ALL        = 0xfffc
	OFPP_CONTROLLER = 0xfffd
	OFPP_LOCAL      = 0xfffe
	OFPP_NONE       = 0xffff
)
