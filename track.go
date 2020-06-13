package deliverytracker

import "time"

type TrackFrom struct {
	Name string     `json:"name"`
	Time *time.Time `json:"time"`
}

type TrackTo struct {
	Name string     `json:"name"`
	Time *time.Time `json:"time"`
}

type TrackState struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TrackProgressStatus struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TrackProgressLocation struct {
	Name string `json:"name"`
}

type TrackProgress struct {
	Time        *time.Time             `json:"time"`
	Status      *TrackProgressStatus   `json:"status"`
	Location    *TrackProgressLocation `json:"location"`
	Description string                 `json:"description"`
}

type Track struct {
	From       *TrackFrom       `json:"from"`
	To         *TrackTo         `json:"to"`
	State      *TrackState      `json:"state"`
	Progresses []*TrackProgress `json:"progresses"`
	Carrier    *Carrier         `json:"carrier"`
}
