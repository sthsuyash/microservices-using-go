package main

import (
	"context"
	"fmt"

	"github.com/sthsuyash/microservices-using-go/application"
)

func main() {
	app := application.New()

	err := app.Routes(context.TODO())
	if err != nil {
		fmt.Println("failed to start server: %w", err)
	}
}
