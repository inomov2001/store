package main

import (
	"context"
	"store/config"
	"store/inventory"
	"store/server/http"
	"store/server/telegram"
	"store/store"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	cfg := config.Load()

	i, err := inventory.NewPostgresInventory(ctx, cfg)
	if err != nil {
		panic(err)
	}

	s := store.New(i)

	httpServer := http.NewServer(s)
	telegramBotServer, err := telegram.NewServer(cfg.TelegramApiToken, s)
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err = httpServer.Run(cfg.StoreHost + cfg.HTTPPort); err != nil {
			panic(err)
		}
	}()

	go func() {
		defer wg.Done()
		telegramBotServer.Run()
	}()

	wg.Wait()
}
