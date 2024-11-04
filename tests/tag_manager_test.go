package synapse_test

import (
	synapse "synapse/src"
	"testing"
)

func TestTagManager(t *testing.T) {
    tm := synapse.NewTagManager()

    // Test tag creation
    tag := tm.NewTag("source1")
    if tag == "" {
        t.Error("Expected non-empty tag")
    }

    // Test tag processing status
    if !tm.IsProcessed(tag) {
        t.Error("Expected tag to be marked as processed")
    }

    // Test pushing new tag
    newTag := "source2-123456"
    tm.PushTag(newTag)

    if !tm.IsProcessed(newTag) {
        t.Error("Expected pushed tag to be marked as processed")
    }
}
