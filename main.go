package main

import (
	"fmt"
	"log"

	"github.com/Tata-Matata/fake-aria-api/deploy"
)

func main() {
	logger := Logger{}
	logger.Initialize()
	fmt.Println("Writing logs to %v", logger.Dir)
	defer logger.Close()

	deployApi, _ := deploy.NewDeployAPI(deploy.DEPLOY_JSON)

	app := App{DeployApi: *deployApi}
	app.Initialize()

	log.Println("Application initialized. Starting...")
	app.Run(":10000")

}
