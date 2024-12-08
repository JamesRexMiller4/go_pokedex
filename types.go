package main

type CliCommand struct {
	Name        string
	Description string
	Callback    func(config *Config) error
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Config struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []Location `json:"results"`
}
