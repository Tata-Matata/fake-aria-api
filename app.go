package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Tata-Matata/fake-aria-api/deploy"
	"github.com/gorilla/mux"
)

type App struct {
	Router    *mux.Router
	DeployApi deploy.DeployAPI
}

func checkErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func (app *App) Initialize() error {

	app.Router = mux.NewRouter().StrictSlash(true)
	app.HandleRoutes()
	return nil
}

func sendJSON(response http.ResponseWriter, data interface{}, statusCode int) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	json.NewEncoder(response).Encode(data)
}

func sendError(response http.ResponseWriter, status int, err string) {
	message := map[string]string{"error": err}

	sendJSON(response, message, status)

}

func (app *App) getDeployments(respWriter http.ResponseWriter, req *http.Request) {
	deploys, err := app.DeployApi.GetRandom()
	if err != nil {
		sendError(respWriter, http.StatusInternalServerError, err.Error())
		return
	}

	sendJSON(respWriter, deploys, http.StatusOK)
}

func (app *App) getDeployment(respWriter http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	deployment, err := app.DeployApi.GetByID(id)

	if err != nil {
		sendError(respWriter, http.StatusBadRequest, "invalid deployment ID")
		return
	}

	sendJSON(respWriter, deployment, http.StatusOK)
}

func (app *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, app.Router))
}

func (app *App) HandleRoutes() {
	app.Router.HandleFunc("/deployments", app.getDeployments).Methods("GET")
	app.Router.HandleFunc("/deployments/{id}", app.getDeployment).Methods("GET")
	//app.Router.HandleFunc("/product", app.createProduct).Methods("POST")
}
