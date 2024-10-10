package main

import (
	"context"
	"file-server/cmd/rest"
	"file-server/lib/config"
	"log"
	"sync"
)

func main() {
	log.Println("Starting file server")

	cfg := config.GetConfig("./cmd", "config", "yaml")

	log.Println("File server started")

	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		ctx := context.Background()

		requestHandler := rest.NewHandler(*cfg)

		err := rest.Run(ctx, cfg, requestHandler)
		if err != nil {
			log.Fatal(err)
		}

		wg.Done()
	}()

	wg.Wait()

}
