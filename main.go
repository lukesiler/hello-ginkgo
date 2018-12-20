package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/lukesiler/hello-ginkgo/api"
)

func main() {
	logger := log.New(os.Stdout, "", 0)
	httpServer, gracefulStopChan, err := api.ServeAPI()
	if err != nil {
		logger.Panic(err)
	}

	sig := <-gracefulStopChan

	logger.Println(fmt.Sprintf("Captured %v. Shutting down HTTP server.", sig))
	httpServer.Shutdown(context.Background())
	logger.Println("Server gracefully shutdown")
}
