package transport

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
)

// Encode Message to []byte
//
// Although sending messages is simple, how to ensure that the server reads the
// data correctly is a problem.
//
// My idea is to calculate the length of the data message (byte) before sending
// data each time. This is called packet size!
//
// After the length is obtained, it cannot be directly sent (because it is an int
// type). The value must be converted to []byte for network transmission, So you
// need a way to convert this value to byte type.
//
// There is an interface under the standard library binary.ByteOrder, and two concrete
// instances (BigEndian and LittleEndian) have been provided for us to use.
//
// binary.BigEndian and binary.LittleEndian refers to the computer bit storage order,
// This is introduced on Wikipedia:
//
//   https://en.wikipedia.org/wiki/Endianness
//
// So we can use one of them to achieve our goal, Use binary.BigEndian here.
//
// The specific operation is to calculate the length of the message body and use
// binary.BigEndian converts it to a byte representation And write directly to buf.
//
// When the data length is successfully written, the real message is written to the
// buf.
//
// Because the uint32 size is 4byte, this part of data can be fixed, That is to say,
// the data we actually returned by this method is actually the following:
//
//    +---------------+---------------------------------------+
//    |    4byte      |             dataLen byte              |
//    +---------------+---------------------------------------+
//    | Packet Length |           Real data message           |
//    +---------------+---------------------------------------+
//
// This ensures that the server first obtains the specific message length when obtaining
// data, and then obtains the real data content according to the specific length.
//
func Encode(message *Message) ([]byte, error) {

	data, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)

	length := uint32(len(data))

	if err := binary.Write(buf, binary.BigEndian, length); err != nil {
		return nil, err
	}

	if err := binary.Write(buf, binary.BigEndian, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Decode buf(read form bufio.Reader) to Message
//
// When get the network packet size of the transport before getting the real transport
// sent by the client.
//
// Before sending a transport each time, the client will first obtain the data stream
// size(use uint32 type), and then use the binary.BigEndian value to encode it into the
// corresponding []byte data type. This is called packet size!
//
// uint32 The value is 4 bytes, So the first thing we receive is actually this packetN.
// After that, we start to get the data size in the buf buffer recN. And compare whether
// recN is equal to the actual data size (packetN+1), If the two are equal, the data will
// be received normally.
//
// The actual data form of buf buffer should be as follows:
//
//    +---------------+---------------------------------------+
//    |    4byte      |              Packet Length            |
//    +---------------+---------------------------------------+
//    | Packet Length |        Real data message length       |
//    +---------------+---------------------------------------+
//
// uint32 Maximum 4 GB data.
//
func Decode(r *bufio.Reader) (*Message, error) {

	// Read the first 4 bytes of data
	packetN, _ := r.Peek(4)
	buf := bytes.NewBuffer(packetN)

	var length uint32
	if err := binary.Read(buf, binary.BigEndian, &length); err != nil {
		return nil, err
	}

	// Buffered returns the number of bytes that are currently readable in the buffer
	recN := r.Buffered()
	if uint32(recN) < length+4 {
		return nil, errors.New("the size of the data in buf does not match the actual data")
	}

	pack := make([]byte, length+4)
	if _, err := r.Read(pack); err != nil {
		return nil, err
	}

	message := new(Message)
	if err := json.Unmarshal(pack[4:], message); err != nil {
		return nil, err
	}

	return message, nil
}
