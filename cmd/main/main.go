package main

import (
	"auth/internal/config"
	"auth/internal/logging"
	"auth/internal/mongodb"
	"auth/internal/web"
	"math/rand"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
)

func init() {
	// get random seed
	rand.Seed(time.Now().UnixNano())
}

func main() {
	cfg := config.InitConfiguration()
	logger := logging.NewLogger("auth.log")
	logger.Info("Connecting to MongoDB")
	mongoClient := mongodb.NewMongoDB(logger, &cfg.MongoDB)
	defer mongoClient.Release()
	err := mongoClient.Connect()
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}
	logger.Info("Connected to MongoDB")
	wApp := web.NewWebServer(logger, &cfg.Web)
	go wApp.Run()

	//
	// Graceful Shutdown
	//

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt, os.Kill)
	takeSig := <-sigChan
	logger.Info("Received terminate, graceful shutdown", zap.String("signal", takeSig.String()))
}
