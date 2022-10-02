package config

type Config struct {
	Version string `json:"version"`
	Salt string `json:"salt"`
	Sites []Site `json:"sites"`
}

type Site struct {
	Domain string `json:"domain"`
	Rotations int `json:"rotations"`
}