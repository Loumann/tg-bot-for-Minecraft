package models

type Player struct {
	Name              string `json:"name"`
	AwaitingInputType bool   `json:"awaiting"`
}

type PlayerLocation struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}
