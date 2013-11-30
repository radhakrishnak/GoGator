package ofp10

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

type OFMatch struct {
	Wildcards                            uint32           /* Wildcard fields. */
	InPort                               uint16           /* Input switch port. */
	DataLayerSource                      net.HardwareAddr //[ETH_ALEN]uint8 /* Ethernet source address. */
	DataLayerDestination                 net.HardwareAddr //[ETH_ALEN]uint8 /* Ethernet destination address. */
	DataLayerVirtualLan                  uint16           /* Input VLAN id. */
	DataLayerVirtualLanPriorityCodePoint uint8            /* Input VLAN priority. */
	Pad                                  [1]uint8         /* Align to 64-bits */
	DataLayerType                        uint16           /* Ethernet frame type. */
	NetworkTypeOfService                 uint8            /* IP ToS (actually DSCP field, 6 bits). */
	NetworkProtocol                      uint8            /* IP protocol or lower 8 bits of ARP opcode. */
	Pad1                                 [2]uint8         /* Align to 64-bits */
	NetworkSource                        net.IP           /* IP source address. */
	NetworkDestination                   net.IP           /* IP destination address. */
	TransportSource                      uint16           /* TCP/UDP source port. */
	TransportDestination                 uint16           /* TCP/UDP destination port. */
}

func NewMatch() *OFMatch {
	m := new(OFMatch)
	m.Wildcards = 0xffffffff
	//	m.DLSrc = make([]byte, ETH_ALEN)
	//	m.DLDst = make([]byte, ETH_ALEN)
	m.DataLayerSource = make([]byte, 6)
	m.DataLayerDestination = make([]byte, 6)

	m.NetworkSource = make([]byte, 4)
	m.NetworkDestination = make([]byte, 4)
	return m
}

/*func LoadFromPacket(PacketData networkUtils.Ethernet, InputPort uint16) *OFMatch 	{
    short scratch;
    int transportOffset = 34;
    ByteBuffer packetDataBB = ByteBuffer.wrap(packetData);
    int limit = packetDataBB.limit();

	m := new(OFMatch)
    m.Wildcards = 0; // all fields have explicit entries

    m.InPort = InputPort;

    if InputPort == ofProtocol.OFPP_ALL	{
		m.Wildcards |= OFPFW_IN_PORT
	}

    // dl dst
    m.DataLayerDestination = make([]byte, 6)
	DataLayerDestination = PacketData.HWDst.String()
    // dl src
   	m.DataLayerSource = make([]byte, 6)
	m.DataLayerSource = eth.HWSrc.String()

    // dl type
    m.DataLayerType = eth.Ethertype

    if m.DataLayerType != 0x8100(unit16) 	{ // need cast to avoid signed
        // bug
        m.DataLayerVirtualLan = 0xffff(uint16);
        m.DataLayerVirtualLanPriorityCodePoint = 0(unit8);
    } else {
        // has vlan tag
        scratch = packetDataBB.getShort();
       	m.DataLayerVirtualLan((unit16) = (0xfff & PacketData.));
        m.DataLayerVirtualLanPriorityCodePoint((byte) = ((0xe000 & scratch) >> 13));
        this.dataLayerType = packetDataBB.getShort();
    }

    switch (m.DataLayerType()) {
    case 0x0800:
        // ipv4
        // check packet length
        scratch = packetDataBB.get();
        scratch = (short) (0xf & scratch);
        transportOffset = (packetDataBB.position() - 1) + (scratch * 4);
        // nw tos (dscp)
        scratch = packetDataBB.get();
        setNetworkTypeOfService((byte) ((0xfc & scratch) >> 2));
        // nw protocol
        packetDataBB.position(packetDataBB.position() + 7);
        this.networkProtocol = packetDataBB.get();
        // nw src
        packetDataBB.position(packetDataBB.position() + 2);
        this.networkSource = packetDataBB.getInt();
        // nw dst
        this.networkDestination = packetDataBB.getInt();
        packetDataBB.position(transportOffset);
        break;
    case 0x0806:
        // arp
        int arpPos = packetDataBB.position();
        // opcode
        scratch = packetDataBB.getShort(arpPos + 6);
        setNetworkProtocol((byte) (0xff & scratch));

        scratch = packetDataBB.getShort(arpPos + 2);
        // if ipv4 and addr len is 4
        if (scratch == 0x800 && packetDataBB.get(arpPos + 5) == 4) {
            // nw src
            this.networkSource = packetDataBB.getInt(arpPos + 14);
            // nw dst
            this.networkDestination = packetDataBB.getInt(arpPos + 24);
        } else {
            setNetworkSource(0);
            setNetworkDestination(0);
        }
        break;
    default:
        setNetworkTypeOfService((byte) 0);
        setNetworkProtocol((byte) 0);
        setNetworkSource(0);
        setNetworkDestination(0);
        break;
    }

    switch (getNetworkProtocol()) {
    case 0x01:
        // icmp
        // type
        this.transportSource = U8.f(packetDataBB.get());
        // code
        this.transportDestination = U8.f(packetDataBB.get());
        break;
    case 0x06:
        // tcp
        // tcp src
        this.transportSource = packetDataBB.getShort();
        // tcp dest
        this.transportDestination = packetDataBB.getShort();
        break;
    case 0x11:
        // udp
        // udp src
        this.transportSource = packetDataBB.getShort();
        // udp dest
        this.transportDestination = packetDataBB.getShort();
        break;
    default:
        setTransportDestination((short) 0);
        setTransportSource((short) 0);
        break;
    }
    return this;
}*/

func (m *OFMatch) Read(b []byte) (n int, err error) {
	// Any non-zero value fields should not be wildcarded.
	if m.InPort != 0 {
		m.Wildcards = m.Wildcards ^ FW_IN_PORT
	}
	if m.DataLayerSource.String() != "00:00:00:00:00:00" {
		m.Wildcards = m.Wildcards ^ FW_DL_SRC
	}
	if m.DataLayerDestination.String() != "00:00:00:00:00:00" {
		m.Wildcards = m.Wildcards ^ FW_DL_DST
	}
	if m.DataLayerVirtualLan != 0 {
		m.Wildcards = m.Wildcards ^ FW_DL_VLAN
	}
	if m.DataLayerVirtualLanPriorityCodePoint != 0 {
		m.Wildcards = m.Wildcards ^ FW_DL_VLAN_PCP
	}
	if m.DataLayerType != 0 {
		m.Wildcards = m.Wildcards ^ FW_DL_TYPE
	}
	if m.NetworkTypeOfService != 0 {
		m.Wildcards = m.Wildcards ^ FW_NW_TOS
	}
	if m.NetworkProtocol != 0 {
		m.Wildcards = m.Wildcards ^ FW_NW_PROTO
	}
	if m.NetworkSource.String() != "0.0.0.0" {
		m.Wildcards = m.Wildcards ^ FW_NW_SRC_ALL
	}
	if m.NetworkDestination.String() != "0.0.0.0" {
		m.Wildcards = m.Wildcards ^ FW_NW_DST_ALL
	}
	if m.TransportSource != 0 {
		m.Wildcards = m.Wildcards ^ FW_TP_SRC
	}
	if m.TransportDestination != 0 {
		m.Wildcards = m.Wildcards ^ FW_TP_DST
	}
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, m.Wildcards)
	//	fmt.Println("Match Read: wildcards ", buf.Bytes())
	binary.Write(buf, binary.BigEndian, m.InPort)
	binary.Write(buf, binary.BigEndian, m.DataLayerSource)
	binary.Write(buf, binary.BigEndian, m.DataLayerDestination)
	binary.Write(buf, binary.BigEndian, m.DataLayerVirtualLan)
	binary.Write(buf, binary.BigEndian, m.DataLayerVirtualLanPriorityCodePoint)
	binary.Write(buf, binary.BigEndian, m.Pad)
	binary.Write(buf, binary.BigEndian, m.DataLayerType)
	binary.Write(buf, binary.BigEndian, m.NetworkTypeOfService)
	binary.Write(buf, binary.BigEndian, m.NetworkProtocol)
	binary.Write(buf, binary.BigEndian, m.Pad1)
	binary.Write(buf, binary.BigEndian, m.NetworkSource)
	binary.Write(buf, binary.BigEndian, m.NetworkDestination)
	binary.Write(buf, binary.BigEndian, m.TransportSource)
	binary.Write(buf, binary.BigEndian, m.TransportDestination)
	//	fmt.Println("Match Read: all ", buf.Bytes())
	n, err = buf.Read(b)
	if n == 0 {
		return
	}
	return n, io.EOF
}

// ofp_flow_wildcards 1.0
const (
	FW_IN_PORT  = 1 << 0
	FW_DL_VLAN  = 1 << 1
	FW_DL_SRC   = 1 << 2
	FW_DL_DST   = 1 << 3
	FW_DL_TYPE  = 1 << 4
	FW_NW_PROTO = 1 << 5
	FW_TP_SRC   = 1 << 6
	FW_TP_DST   = 1 << 7

	FW_NW_SRC_SHIFT = 8
	FW_NW_SRC_BITS  = 6
	FW_NW_SRC_MASK  = ((1 << FW_NW_SRC_BITS) - 1) << FW_NW_SRC_SHIFT
	FW_NW_SRC_ALL   = 32 << FW_NW_SRC_SHIFT

	FW_NW_DST_SHIFT = 14
	FW_NW_DST_BITS  = 6
	FW_NW_DST_MASK  = ((1 << FW_NW_DST_BITS) - 1) << FW_NW_DST_SHIFT
	FW_NW_DST_ALL   = 32 << FW_NW_DST_SHIFT

	FW_DL_VLAN_PCP = 1 << 20
	FW_NW_TOS      = 1 << 21

	FW_ALL = ((1 << 22) - 1)
)
