package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/konda-manish/internal/config"
	"github.com/konda-manish/internal/http/handlers/student"
	"github.com/konda-manish/internal/storage/sqlite"
)

func main() {

	// print("Hello world")
	// fmt.Print("Welcome to the Students API")
	// Load config
	cfg := config.MustLoad()
	// fmt.Printf("config: %+v\n", cfg)

	//database setup
	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("storage_path", cfg.StoragePath))

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New(storage))         //this is to handle the new student creation
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage)) //this is to handle the new student creation
	router.HandleFunc("GET /api/students", student.GetList(storage))      //this is to handle the new student creation
	router.HandleFunc("DELETE /api/students/{id}", student.DeleteById(storage))

	//set up server

	server := &http.Server{ // this is to create a new server
		Addr:    cfg.HTTPServer.Address,
		Handler: router,
	}

	fmt.Println("server started on port", cfg.HTTPServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {

		err := server.ListenAndServe() // this is to start the server // this is to listen and serve the server

		if err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done

	slog.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown successfully")

}
