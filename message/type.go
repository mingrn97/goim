package message

type MessageType string

const (
	PlainType MessageType = "plaintext"
	ImageType MessageType = "image"
	VideoType MessageType = "video"
	MediaType MessageType = "media"
)
