package main

import (
	"fmt"
	"log"
	"net/http"
	"test/infrastructure"
	"test/restserver/handlers"
	"test/usecases"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// infrastructure
	users := infrastructure.NewUserRepository()
	products := infrastructure.NewProductRepository()
	orders := infrastructure.NewOrderRepository()
	// infrastructure

	useCases := usecases.New(users, products, orders)

	r := mux.NewRouter()
	// Middleware: Logging
	r.Use(loggingMiddleware)
	// Middleware: Recovery
	r.Use(recoveryMiddleware)
	// Define routes
	r.HandleFunc("/users", handlers.NewRegisterUser(useCases).ServeHTTP).Methods(http.MethodPost)
	r.HandleFunc("/order", handlers.NewMakeOrder(useCases).ServeHTTP).Methods(http.MethodPost)

	// Start the server
	port := 8080
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server listening on port %d...\n", port)
	log.Fatal(server.ListenAndServe())
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		log.Printf("[%s] %s %s %v", r.Method, r.RequestURI, r.Proto, time.Since(startTime))
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
