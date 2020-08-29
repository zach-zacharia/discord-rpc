package discordrpc

type handshake struct {
	V        string `json:"v"`
	ClientID string `json:"client_id"`
}

// Data section of the RPC response
type Data struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Args seems to contain the most data, Pid here is pandatory
type Args struct {
	Pid      int       `json:"pid"`
	Activity *Activity `json:"activity,omitempty"`
}
