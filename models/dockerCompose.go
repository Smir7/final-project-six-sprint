package models

// DockerCompose is a main object in Docker-Compose file
type DockerCompose struct {
	Version  string   `json:"version" yaml:"version"`
	Services Services `json:"services" yaml:"services"`
}

// Web is a service in Docker-Compose file
type Web struct {
	Build   string   `json:"build" yaml:"build"`
	Ports   []string `json:"ports" yaml:"ports"`
	Volumes []string `json:"volumes" yaml:"volumes"`
	Links   []string `json:"links" yaml:"links"`
}

// Database is a service in Docker-Compose file
type Database struct {
	Image       string   `json:"image" yaml:"image"`
	Environment []string `json:"environment" yaml:"environment"`
	Volumes     []string `json:"volumes" yaml:"volumes"`
}

// Services is a services run in Docker
type Services struct {
	Web      Web      `json:"web" yaml:"web"`
	Database Database `json:"database" yaml:"database"`
}
