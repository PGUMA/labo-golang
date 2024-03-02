package main

import (
	"fmt"
	// Support for structured logs
	"log/slog"
	"os"
)

func init() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))
}

func main() {
	fmt.Println("Hello Golang Labo")
	slog.Info("Hello Golang Labo", slog.String("event", "sample logging"))
}
