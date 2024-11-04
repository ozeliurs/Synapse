package synapse_test

import (
	"bytes"
	synapse "synapse/src"
	"testing"
)

func TestNewMessage(t *testing.T) {
    msg := synapse.NewMessage("TEST", "testKey", "testValue", "source", "target")

    if msg.Type != "TEST" {
        t.Errorf("Expected Type 'TEST', got %s", msg.Type)
    }
    if msg.Key != "testKey" {
        t.Errorf("Expected Key 'testKey', got %s", msg.Key)
    }
    if msg.Value != "testValue" {
        t.Errorf("Expected Value 'testValue', got %v", msg.Value)
    }
    if msg.TTL != 10 {
        t.Errorf("Expected TTL 10, got %d", msg.TTL)
    }
}

func TestMessageEncodeDecode(t *testing.T) {
    originalMsg := synapse.NewMessage("TEST", "testKey", "testValue", "source", "target")

    var buf bytes.Buffer
    err := originalMsg.Encode(&buf)
    if err != nil {
        t.Fatalf("Failed to encode message: %v", err)
    }

    decodedMsg := &synapse.Message{}
    err = decodedMsg.Decode(&buf)
    if err != nil {
        t.Fatalf("Failed to decode message: %v", err)
    }

    if originalMsg.Type != decodedMsg.Type {
        t.Errorf("Message type mismatch: expected %s, got %s", originalMsg.Type, decodedMsg.Type)
    }
}
