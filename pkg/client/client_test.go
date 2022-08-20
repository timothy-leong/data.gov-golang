package client

import (
	"fmt"
	"testing"
	"time"
)

func TestCarparkAvailability(t *testing.T) {
	client := NewDataGovClient()
	_, err := client.CarparkAvailability(time.Now())

	if err != nil {
		t.Error(err)
	}

	// Check that this result has been cached
	if len(client.cache) != 1 {
		t.Errorf("Client cache is empty but the result should have been cached")
	}
}
