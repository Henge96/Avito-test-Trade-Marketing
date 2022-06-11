package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"

	"depends/internal/api"
	"depends/internal/app"
	"depends/internal/repo"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal(err)
	}
	cfg := repo.Config{
		Driver:   viper.GetString("db.driver"),
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.name"),
		SSLMode:  viper.GetString("db.sslmode"),
	}

	err = run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func run(cfg repo.Config) error {

	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return fmt.Errorf("sqlx.Connect: %w", err)
	}
	defer db.Close()
	rep := repo.New(db)

	ap := app.New(rep)

	server := &http.Server{
		Addr:    viper.GetString("server.host") + viper.GetString("server.port"),
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

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(viper.GetInt("server.timeout.server_timeout"))*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return nil

}

func initConfig() error {
	viper.AddConfigPath("cmd")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
