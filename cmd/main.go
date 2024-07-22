package main

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/KillReall666/schooldocumentmanagment/internal/config"
	"github.com/KillReall666/schooldocumentmanagment/internal/handlers/create"
	"github.com/KillReall666/schooldocumentmanagment/internal/handlers/read"
	"github.com/KillReall666/schooldocumentmanagment/internal/handlers/readall"
	"github.com/KillReall666/schooldocumentmanagment/internal/handlers/update"
	"github.com/KillReall666/schooldocumentmanagment/internal/service"
	"github.com/KillReall666/schooldocumentmanagment/internal/storage/postgres"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	cfg, err := config.New()
	if err != nil {
		panic(err)
	}

	store, err := postgres.New(cfg.DBPath)
	if err != nil {
		log.Fatal("db not connected", err)
	}
	log.Println("db connected")

	serv := service.New(cfg, store)

	httpServer := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/create", create.NewCreateHandler(serv).Create)
	http.HandleFunc("/read", read.NewReadHandler(serv).Read)
	http.HandleFunc("/update", update.NewUpdateHandler(serv).Update)
	http.HandleFunc("/readall", readall.NewAllPublicationsHandler(serv).ReadAll)

	log.Println("Starting server on", httpServer.Addr)

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		return httpServer.Shutdown(context.Background())
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \n", err)
	}
}
