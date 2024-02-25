package act

// Parameter is a parameter that can be passed to an action by the system (runner),
// like a logger, a database connection, or any other data.
type Parameter struct {
	Key   string
	Value any
}

// ActionFunc is a function that can be used to run an action. Optionally, it can
// receive parameters like a logger, a database connection, or any other data that
// the action needs to run.
type ActionFunc func(data map[string]any, params ...Parameter) (any, error)

// Action is any action that can be performed by the system.
type Action struct {
	Name     string         `json:"name"`
	Metadata map[string]any `json:"metadata"`
	Sync     bool           `json:"sync"`
	Terminal bool           `json:"terminal"`
	Run      ActionFunc     `json:"-"`
}
