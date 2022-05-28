package transport

type Message struct {

	// message type
	Type MessageType `json:"type"`

	// message data
	Data string `json:"data"`
}
