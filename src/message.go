package synapse

import (
	"encoding/gob"
	"io"
)

type Message struct {
	Type   string
	Key    string
	Value  interface{}
	Source string
	Target string
	Tag    string
	TTL    int
	MRR    int
}

func NewMessage(msgType string, key string, value interface{}, source string, target string) *Message {
	return &Message{
		Type:   msgType,
		Key:    key,
		Value:  value,
		Source: source,
		Target: target,
		TTL:    10,
		MRR:    3,
	}
}

func (m *Message) Encode(w io.Writer) error {
	return gob.NewEncoder(w).Encode(m)
}

func (m *Message) Decode(r io.Reader) error {
	return gob.NewDecoder(r).Decode(m)
}
