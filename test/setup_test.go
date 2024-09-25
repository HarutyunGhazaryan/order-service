package test

import (
	"OrderService/internal/app"
	"OrderService/internal/cache"
	"OrderService/internal/config"
	db "OrderService/internal/generated"
	"OrderService/internal/models"
	"OrderService/utils"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	serverStarted bool
	serverCtx     context.Context
	serverCancel  context.CancelFunc
	dbQueries     db.Queries
	testOrder     models.Order
	cfg           *config.Config
)

func TestMain(m *testing.M) {
	err := os.Chdir("..")
	if err != nil {
		log.Fatalf("Could not change directory: %v", err)
	}

	cfg = config.LoadConfig()

	dbConn := app.ConnectToDB(cfg.DB_URL)
	defer dbConn.Close()

	dbQueries = *db.New(dbConn)

	if !isServerRunning() {
		startServer()
	}

	// var err error
	testOrder, err = utils.CreateTestOrder(context.Background(), &dbQueries)
	if err != nil {
		fmt.Printf("Failed to set up test order: %v\n", err)
		stopServer()
		os.Exit(1)
	}

	code := m.Run()

	if err := utils.DeleteTestOrder(context.Background(), &dbQueries, testOrder.OrderUID); err != nil {
		fmt.Printf("Failed to delete test order: %v\n", err)
	}

	if serverStarted {
		stopServer()
	}

	os.Exit(code)
}

func isServerRunning() bool {
	URL := fmt.Sprintf("http://localhost:%s/health", cfg.PORT)
	resp, err := http.Get(URL)

	if err != nil {
		return false
	}
	resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

func startServer() {
	serverCtx, serverCancel = context.WithCancel(context.Background())
	cfg := config.LoadConfig()
	orderCache := cache.NewOrderCache()

	go app.StartHTTPServer(serverCtx, cfg.PORT, orderCache, &dbQueries)

	time.Sleep(2 * time.Second)
	serverStarted = true
}

func stopServer() {
	if serverStarted {
		serverCancel()
	}
}
