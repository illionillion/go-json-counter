package utils

type NameCount struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Counter struct {
	Count int         `json:"count"`
	Data  []NameCount `json:"data"`
}
