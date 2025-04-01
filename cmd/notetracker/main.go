package main

import "github.com/MikVG/note-tracker/internal/config"

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		panic(err)
	}
}
