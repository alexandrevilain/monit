package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"

	"github.com/alexandrevilain/monit/config"
	"github.com/alexandrevilain/monit/handler"
	"github.com/alexandrevilain/monit/job"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	configPath := flag.String("config", "config.yaml", "Path to the config file")
	flag.Parse()

	config, err := config.LoadConfigFromFile(*configPath)
	if err != nil {
		panic(err)
	}

	serviceStateStore := job.NewServicesStateStore()
	jh := job.NewHandler()
	err = jh.CreateJobsFromServices(serviceStateStore, config.Services)
	if err != nil {
		log.Fatalf(err.Error())
	}

	srv := &http.Server{
		Addr:    config.HTTP.ListenAddress,
		Handler: handler.New(serviceStateStore, &config.HTTP),
	}

	// Start jobs:
	jh.Start()

	// Check for a closing signal
	go func() {
		sigquit := make(chan os.Signal, 1)
		signal.Notify(sigquit, os.Interrupt, os.Kill)

		sig := <-sigquit
		if err := jh.Stop(); err != nil {
			log.Printf("Unable to stop jobs")
		}
		log.Printf("Gracefully shutting down server, caught: %+v", sig)
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("Unable to shut down server: %v", err)
		} else {
			log.Println("Server stopped")
		}
	}()

	// Start server
	log.Printf("Starting HTTP Server. Listening at %q", srv.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	}

}
