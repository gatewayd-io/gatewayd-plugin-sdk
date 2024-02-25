package act

type Signal struct {
	Name     string         `json:"name"`
	Metadata map[string]any `json:"metadata"`
}

func (s *Signal) ToMap() map[string]any {
	return map[string]any{
		Name:     s.Name,
		Metadata: s.Metadata,
	}
}

func Terminate() *Signal {
	return &Signal{
		Name: "terminate",
		Metadata: map[string]any{
			"terminate": true,
		},
	}
}

func Log(level, message string, fields map[string]any) *Signal {
	metadata := map[string]any{
		"log":     true,
		"level":   level,
		"message": message,
	}

	if fields != nil {
		for k, v := range fields {
			if k == "level" || k == "message" {
				continue
			}
			metadata[k] = v
		}
	}

	return &Signal{
		Name:     "log",
		Metadata: metadata,
	}
}
