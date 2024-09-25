package app

import (
	"OrderService/internal/cache"
	db "OrderService/internal/generated"
	"OrderService/internal/handler"
	"context"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// StartHTTPServer starts an HTTP server on the specified port and uses the provided context to manage server shutdown.
// Server shutdown via context is intended only for testing purposes.

func StartHTTPServer(ctx context.Context, port string, orderCache *cache.OrderCache, dbQueries *db.Queries) *http.Server {
	r := chi.NewRouter()
	r.Get("/order/{order_uid}", handler.OrderHandler(orderCache))
	r.Get("/display/{order_uid}", handler.DisplayHandler(orderCache))
	r.Get("/display", handler.DisplayHandler(orderCache))
	r.Get("/db-order/{order_uid}", handler.DBOrderHandler(dbQueries))
	r.Get("/health", handler.HealthHandler)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("internal/static"))))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		if err := server.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down the server: %v", err)
		}
	}()

	log.Printf("Starting server on %s", server.Addr)
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("ListenAndServe failed: %v", err)
	}

	return server
}
