package common

// Config represents the entire configuration from JSON file
type Config struct {
	URLPrefix string    `json:"urlPrefix"`
	Token     string    `json:"token"`
	Services  []Service `json:"services"`
}

// Service represents a single service configuration
type Service struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Param  string `json:"param"`
	Enable bool   `json:"enable"`
}
