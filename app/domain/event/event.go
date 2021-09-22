package event

// Event is the domain model
// for the event date type
type Event struct {
	ID       uint64 `json:"id,omitempty"`
	Type     string `json:"type"`
	Time     string `json:"time"`
	Identity string `json:"identity"`
}
