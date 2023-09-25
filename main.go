package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/Clarks0n/go-orders-api/application"
)

func main() {
	app := application.New()

	// Shutdown server if no connection, 1.set Context
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) 
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("Failed to start app:", err)
	}
}
