package main

import (
	"context"
	"fmt"

	"github.com/Abdullah-Hamad-Almousa/server-chi/application"
)

func main() {

	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("Failed to start app: ", err)
	}
}
