package main

import (
	"sync"

	"go.uber.org/zap"
)

var (
	log zap.SugaredLogger
	wg  sync.WaitGroup
)

func init() {

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any

	log = *logger.Sugar()
	log.Info("Initiated logging with zap")

	// init config
	ReadConfigFile()
}

func main() {
	log.Info("Starting pulsar client..")

	NewPulsarClient()

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		SendMessageAsProducer("mytopic", "hello")
		wg.Done()
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		Consume("mytopic")
		wg.Done()
	}(&wg)

	wg.Wait()
	log.Info("Finished")
}
