package douyu

import (
	"encoding/binary"
	"log"
)

var msgType = make([]byte, 2)

func init() {
	binary.LittleEndian.PutUint16(msgType, uint16(0x2b1))
}

// Prepare the message to send to danmu server
//
// The message consists of the following parts:
// - [4]byte{} which is the length of message
// - [4]byte{} identical to the above part
// - [2]byte{} message type. always 0x2b1 for request to server (0x2b2 for resp from server)
// - '0x00' reserved bit for encryption method.
// - '0x00' reserved bit
// - [n]byte{} the message (content)
// - '0x00' the end of message '\0'
//
// Therefore total length would be 3 * 4 + n + 1
func Dumps(content []byte) []byte {
	var buf = make([]byte, 12+len(content)+1)

	// Length
	// seems the length ignores the message type, and 2 reserved bits.
	// therefore length = 2 * 4 + n + 1
	binary.LittleEndian.PutUint32(buf[:4], uint32(len(content)+9))
	copy(buf[4:8], buf[:4])

	// msg type part
	copy(buf[8:10], msgType)

	// encrypt, reserve buf[10:12] are revered bits for now

	// content
	copy(buf[12:], content)

	// "\0"
	// buf[len(buf)-1] = '\x00'

	return buf
}
