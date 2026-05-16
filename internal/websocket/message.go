package websocket

type MessageType string

var (
	TYPE_REGISTER MessageType = "register"
	TYPE_JOB      MessageType = "job"
	TYPE_ACK      MessageType = "ack"
	TYPE_PING     MessageType = "ping"
)

type Message struct {
	Type string `json:"type"`
}
