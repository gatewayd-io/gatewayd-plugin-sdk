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
