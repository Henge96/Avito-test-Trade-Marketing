package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"depends/internal/api"
	"depends/internal/app"
	"depends/internal/config"
	"depends/internal/repo"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	flag.Parse()

	configStruct, err := config.TakeConfigFromYaml(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = run(configStruct)
	if err != nil {
		log.Fatal(err)
	}
}

func run(cfg *config.Config) error {

	db, err := sqlx.Open(cfg.DB.Driver, fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.DB.HostDB, cfg.DB.PortDB, cfg.DB.User, cfg.DB.NameDB, cfg.DB.Password, cfg.DB.SSLMode))
	if err != nil {
		return fmt.Errorf("sqlx.Open: %w", err)
	}
	defer db.Close()
	rep := repo.New(db)

	ap := app.New(rep)

	server := &http.Server{
		Addr:    cfg.Server.Host + ":" + cfg.Server.Port.Http,
		Handler: api.NewRouter(ap),
	}

	go func() {
		err = server.ListenAndServe()
		if err == http.ErrServerClosed {
		} else {
			log.Fatal(err)
		}
	}()

	chanSig := make(chan os.Signal, 1)

	signal.Notify(chanSig, os.Interrupt)

	<-chanSig

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout.ServerTimeout)*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
