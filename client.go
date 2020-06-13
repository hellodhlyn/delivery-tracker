package deliverytracker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Client interface {
	ListCarriers() ([]*Carrier, error)
	FindCarriersByName(name string) ([]*Carrier, error)
	GetTrack(carrierID string, trackID string) (*Track, error)
}

type ClientImpl struct {
	httpClient *http.Client
	endpoint   string
}

// NewClient creates and return a new API client for the delivery-tracker service.
func NewClient() (Client, error) {
	httpClient := &http.Client{}
	return &ClientImpl{
		httpClient: httpClient,
		endpoint:   "https://apis.tracker.delivery",
	}, nil
}

// ListCarriers returns list of all carriers.
func (c *ClientImpl) ListCarriers() ([]*Carrier, error) {
	return c.listCarriers()
}

// FindCarriersByName returns list of carriers that containing given name.
func (c *ClientImpl) FindCarriersByName(name string) ([]*Carrier, error) {
	all, err := c.listCarriers()
	if err != nil {
		return nil, err
	}

	carriers := []*Carrier{}
	for _, carrier := range all {
		if strings.Contains(carrier.Name, name) {
			carriers = append(carriers, carrier)
		}
	}

	return carriers, nil
}

// GetTrack returns a tracking information.
func (c *ClientImpl) GetTrack(carrierID, trackID string) (*Track, error) {
	path := fmt.Sprintf("/carriers/%s/tracks/%s", carrierID, trackID)
	res, err := c.httpClient.Get(c.endpoint + path)
	if err != nil {
		return nil, err
	}

	var track Track
	err = json.NewDecoder(res.Body).Decode(&track)
	if err != nil {
		return nil, err
	}

	return &track, nil
}

func (c *ClientImpl) listCarriers() ([]*Carrier, error) {
	res, err := c.httpClient.Get(c.endpoint + "/carriers")
	if err != nil {
		return nil, err
	}

	var carriers []*Carrier
	err = json.NewDecoder(res.Body).Decode(&carriers)
	if err != nil {
		return nil, err
	}

	return carriers, nil
}
