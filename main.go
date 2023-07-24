package main

import (
	"coffee_tym/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	l := log.New(os.Stdout, "", log.Default().Flags())
	ph := handlers.NewProductsHandler(l)
	sm := mux.NewRouter()

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.PutProducts)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.GetProducts)

	handler := cors.AllowAll().Handler(sm)

	server := http.Server{
		Addr:         "localhost:6969",
		Handler:      handler,
		ErrorLog:     l,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  1e0 * time.Second,
	}

	go func() {
		l.Println("Server started")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
			os.Exit(1)
		}
	}()

	sch := make(chan os.Signal, 1)

	signal.Notify(sch, os.Interrupt)

	sig := <-sch
	l.Println("got ", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Shutdown(ctx)
}
