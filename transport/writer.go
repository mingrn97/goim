package transport

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
)

func WritePlaintext(c net.Conn, text string) {

	var message = &Message{
		Type: PlainType,
		Data: text,
	}

	Write(c, message)
}

func Write(c net.Conn, message *Message) {

	// Although sending messages is simple, how to ensure that the server
	// reads the data correctly is a problem.

	// My idea is to calculate the length of the data message (byte) before
	// sending data each time. This is called packet size!

	// After the length is obtained, it cannot be directly sent (because it
	// is an int type). The value must be converted to []byte for network
	// transmission, So you need a way to convert this value to byte type.

	// There is an interface under the standard library binary.ByteOrder,
	// and two concrete instances (BigEndian and LittleEndian) have been provided
	// for us to use.

	// BigEndian and LittleEndian refers to the computer bit storage order, This is
	// introduced on Wikipedia:
	//
	//   https://en.wikipedia.org/wiki/Endianness
	//

	// So we can use one of them to achieve our goal, Use binary.BigEndian here.

	// The specific operation is to calculate the length of the message body and use
	// binary.BigEndian converts it to a byte representation(variable dataLen). Then
	// send the dataLen to the server. When the data is successfully sent, send the
	// specific data message.

	// This ensures that the server first obtains the specific message length when obtaining
	// data, and then obtains the real data content according to the specific length.

	// Because the uint32 size is 4byte, this part of data can be fixed, That is to say, the
	// data we actually send is as follows:

	//    +---------------+---------------------------------------+
	//    |    4byte      |             dataLen byte              |
	//    +---------------+---------------------------------------+
	//    | Packet Length |           Real data message           |
	//    +---------------+---------------------------------------+

	data, _ := json.Marshal(message)
	fmt.Println(string(data))

	var me = new(Message)
	_ = json.Unmarshal(data, me)

	// data length
	var dataLen [4]byte
	binary.BigEndian.PutUint32(dataLen[:4], uint32(len(data)))

	if n, e := c.Write(dataLen[:4]); n != 4 || e != nil {
		log.Println("send content length fail", e)
	} else {
		// If the data message length is successfully sent,
		// continue to send the actual data!
		_, _ = c.Write(data)
	}
}
