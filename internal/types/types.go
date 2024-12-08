package main

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config) error
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
