package chat

type Message struct {
	Data     Data     `json:"data"`
	Metadata Metadata `json:"metadata"`
}

type Data struct {
	Text string `json:"text"`
}

type Metadata struct {
	Timestamp int64 `json:"timestamp"`
}
