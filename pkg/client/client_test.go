package client

import (
	"testing"
	"time"

	"github.com/timothy-leong/data.gov-golang/pkg/endpoints"
)

func TestCarparkAvailability(t *testing.T) {
	client := NewDataGovClient()
	_, err := client.CarparkAvailability(time.Now())

	if err != nil {
		t.Error(err)
	}

	// Check that this result has been cached
	if _, ok := client.cache[endpoints.CarparkAvailability]; !ok {
		t.Errorf("Client cache is empty but the result should have been cached")
	}
}

func TestTrafficImages(t *testing.T) {
	client := NewDataGovClient()
	_, err := client.TrafficImages(time.Now())

	if err != nil {
		t.Error(err)
	}

	// Check that this result has been cached
	if _, ok := client.cache[endpoints.TrafficImages]; !ok {
		t.Errorf("Client cache is empty but the result should have been cached")
	}
}

func TestLatestUvIndexReadings(t *testing.T) {
	client := NewDataGovClient()
	_, err := client.LatestUVReadings(time.Now())

	if err != nil {
		t.Error(err)
	}

	// Check that this result has been cached
	if _, ok := client.cache[endpoints.UvIndex]; !ok {
		t.Errorf("Client cache is empty but the result should have been cached")
	}
}
