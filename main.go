package main

import (
	"fmt"
	"log"

	"github.com/Tata-Matata/fake-aria-api/deploy"
	ep "github.com/Tata-Matata/fake-aria-api/eventpusher"
	"github.com/Tata-Matata/fake-aria-api/storage"
	"github.com/Tata-Matata/fake-aria-api/util"
)

func main() {
	// Initialize logger
	logger := util.Logger{}
	err := logger.Initialize()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)

	}
	fmt.Printf("Writing logs to %v\n", logger.Dir)
	defer logger.Close()

	// Initialize event pushers
	pushers := []ep.EventPusher{
		ep.DeployEventPusher{},
	}
	exporterURL := EXPORTER_URL + EVENT_EXPORTER_ENDPOINT

	//Without "go" pushers never start because the code after them never runs.
	// pushers run forever and block the main thread.
	go ep.StartEventPushers(pushers, exporterURL)

	// Initialize APIs
	deployApi, _ := deploy.NewDeployAPI(deploy.DEPLOY_JSON)
	storageApi, _ := storage.NewStorageAPI(storage.DEPLOY_JSON)

	app := App{DeployApi: *deployApi, StorageAPI: storageApi}

	// Initialize application
	app.Initialize()

	log.Println("Application initialized. Starting...")
	app.Run(FAKE_API_SERVER_PORT)

}
