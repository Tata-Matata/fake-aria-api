package main

import (
	"fmt"
	"log"

	"github.com/Tata-Matata/fake-aria-api/deploy"
	"github.com/Tata-Matata/fake-aria-api/storage"
)

func main() {
	logger := Logger{}
	logger.Initialize()
	fmt.Printf("Writing logs to %v\n", logger.Dir)
	defer logger.Close()

	deployApi, _ := deploy.NewDeployAPI(deploy.DEPLOY_JSON)
	storageApi, _ := storage.NewStorageAPI(storage.DEPLOY_JSON)

	app := App{DeployApi: *deployApi, StorageAPI: storageApi}
	app.Initialize()

	log.Println("Application initialized. Starting...")
	app.Run(":10000")

}
