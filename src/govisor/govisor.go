package main

import ("fmt"
        "net"
//        "bufio"
        "log"
	"io"
	"os/exec"
	"encoding/binary"
	"bytes"
        )

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

func main() {
	service := ":6633"
	taddr, errr := net.ResolveTCPAddr("tcp", service)
	if errr != nil {
        	fmt.Printf("Failure to listen: %s\n", errr.Error())
        }
	l, err := net.ListenTCP("tcp", taddr)
	if err != nil {
	fmt.Printf("Failure to listen: %s\n", err.Error())
	}

	for {	
		fmt.Printf("Listening...")
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		defer conn.Close()	
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		Echo(conn)
	}
}

func FeaturesRequest() *Header {
	h := NewHeader()
	h.Type = T_FEATURES_REQUEST
	return h
}

func EchoReply() *Header {
        h := NewHeader()
        h.Type = T_ECHO_REPLY
        return h
}
func PacketOut() *Header {
        h := NewHeader()
        h.Type = T_PACKET_OUT
        return h
}

func Send(data interface{},conn net.Conn) {

	buf_Out := new(bytes.Buffer)
	binary.Write(buf_Out, binary.BigEndian, data)
	if _, err := conn.Write(buf_Out.Bytes()); err != nil {
		fmt.Println("Sending msg to failed! ", err)
	}
}
func Echo(c net.Conn) {
	//defer c.Close()
	buf := make([]byte, 1500)
	if _, err := c.Read(buf); err != nil {
		if err != io.EOF {
			fmt.Println("ERROR::Switch.Receive::Read:", err)
		}
	}
	//c.Write(buf)
	var response []byte = SendPacket(buf)
	fmt.Println(response)
	var b []byte
	b = []byte{ 0x01, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00 }
	switch buf[1] {
	case T_HELLO:
		c.Write(b)
		Send(FeaturesRequest(),c)	
		//c.Write(response)
		break
	case T_PACKET_IN:
		Send(PacketOut(),c)
		break
	case T_ECHO_REQUEST:
		Send(EchoReply(),c)
		break
		
	default:
		c.Write(response)
	}
	/*
	fmt.Println(buf)
	line, err := bufio.NewReader(c).ReadString('\n')
	if err != nil {
		fmt.Printf("Failure to read: %s\n", err.Error())
		return
	}
	_, err = c.Write([]byte(line))
	fmt.Printf(line)
	if err != nil {
		fmt.Printf("Failure to write: %s\n", err.Error())
		return
	}*/
	return
}

func SendPacket(buf []byte)(buf1 []byte){
	sconn, err := net.Dial("tcp","127.0.0.1:6635")
	if err != nil {
		fmt.Println("Oops error in sendpacket")
	}
	//buf1 = make([]byte, 1500)
	//if buf[1] == T_HELLO {
	  //      buf1 = []byte{ 0x01, 0x00, 0x00, 0x08, 0x00, 0x00, 0x00, 0x00 }
	//	return
	//}
	_, e := sconn.Write(buf)
	if e != nil {
		fmt.Println("packetout writing error")
	}
	buf1 = make([]byte, 1500)
	//fmt.Println(cmd)	
	cmd.Run()
	if _, e1 := sconn.Read(buf1); e1 != nil {
		fmt.Printf("Failure to read: %s\n", e1.Error())
                return
	}
	//fmt.Println(buf1)
	return
}
