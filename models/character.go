package models

type Character struct {
	Name       string
	Health     int
	Attack     int
	Defense    int
	Experience int
}

var Characters = make(map[string]Character)