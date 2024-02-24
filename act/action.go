package act

import "github.com/rs/zerolog"

// ActionFunc is a function that can be used to run an action.
type ActionFunc func(logger zerolog.Logger, data map[string]any) (any, error)

// Action is any action that can be performed by the system.
type Action struct {
	Name     string         `json:"name"`
	Metadata map[string]any `json:"metadata"`
	Sync     bool           `json:"sync"`
	Terminal bool           `json:"terminal"`
	Run      ActionFunc     `json:"-"`
}
