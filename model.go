package main

type Status struct{
	Name string `json:"name"`
	Health HealthState `json:health_state`
	Reason string `json:"reason_for_failure`
}

type HealthState int

const (
	OK HealthState = iota
	WARN
	ERROR
)

type State interface {
	Value() string
}

func(s HealthState) Value() string {
	switch s {
	case OK: return "OK"
	case WARN: return "WARN"
	case ERROR: return "ERROR"
	}
	return "Not defined"
}
