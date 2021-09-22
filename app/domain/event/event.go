package event

type Event struct {
	ID       uint64 `json:"id,omitempty"`
	Type     string `json:"type"`
	Time     string `json:"time"`
	Identity string `json:"identity"`
}
