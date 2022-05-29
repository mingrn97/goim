package transport

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"net"

	"itumate.com/im/exp"
)

func Read(c net.Conn) (*Message, error) {

	buf := make([]byte, 1024)

	pn, err := c.Read(buf)
	if err == io.EOF {
		return nil, exp.ConnClose
	}
	if err != nil {
		return nil, exp.ReadErr
	}

	// When get the network packet size of the transport before getting the real
	// transport sent by the client.

	// Before sending a transport each time, the client will first obtain the data
	// stream size(use uint32 type), and then use the binary.BigEndian value to encode
	// it into the corresponding []byte data type. This is called packet size!

	// uint32 The value is 4 bytes, So the first thing we receive is actually this
	// packetN. After that, we start to read the connection data circularly until the
	// read data size recN is equal to packetN, indicating that the real data reading
	// is completed.

	// uint32 Maximum 4 GB data.

	// Although the client sends the data in two steps (first send the length of the data
	// message, and then send the real data message after the transmission is successful),
	// the server cannot guarantee that the length of the data message obtained for the first
	// time must be 4.

	// This means that the value of pn may be greater than or equal to 4, This is because the
	// data message read by the server for the first time may have the following two conditions:

	//    +---------------+---------------------------------------+
	//    |    4byte      |      n byte (n <= len(buf - 4))       |
	//    +---------------+---------------------------------------+
	//    | Packet Length |    Partial real data message length   |
	//    +---------------+---------------------------------------+

	// So, If the length of buf is greater than 4, it means that the part exceeding 4 is part of
	// the real data message. The correct approach is to judge whether the length of buf(pn) is
	// greater than 4 in the first step.
	//
	// If it is greater than 4, the greater part will be intercepted and downloaded into packets.

	var recN uint32 = 0
	packetN := binary.BigEndian.Uint32(buf[:4])

	var packets []byte
	if pn > 4 {
		recN = uint32(pn - 4)
		packets = append(packets, buf[4:pn]...)
	}

	for {

		// This is because a previous acquisition has been made, and the size of the buf is 1024.
		// Perhaps the message actually sent by the client does not exceed the length of the buf,
		// and it has been read for the first time.
		if recN >= packetN {
			break
		}

		n, err := c.Read(buf)
		if err != nil {
			break
		}

		recN += uint32(n)

		packets = append(packets, buf[:n]...)
	}

	var message = new(Message)
	_ = json.Unmarshal(packets[:packetN], message)

	return message, nil
}
