package deploy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type DeployAPI struct {
	Deployments []interface{} `json:"deployments"`
}

func (api *DeployAPI) GetRandom() ([]interface{}, error) {

	return api.Deployments, nil
}

func (api *DeployAPI) GetByID(id string) (interface{}, error) {

	for _, item := range api.Deployments {
		m, ok := item.(map[string]interface{})
		if !ok {
			continue
		}
		if v, ok := m["id"]; ok && v == id {
			return item, nil
		}
	}
	log.Printf("deployment with id %v not found", id)
	return nil, fmt.Errorf("deployment with id %v not found", id)
}

func NewDeployAPI() (*DeployAPI, error) {

	emptyApi := &DeployAPI{
		Deployments: []interface{}{},
	}

	data, err := os.ReadFile(DEPLOY_FILE)
	if err != nil {
		log.Printf("failed to read JSON file: %v", err)
		return emptyApi, err
	}

	var deployments []interface{}
	if err := json.Unmarshal(data, &deployments); err != nil {
		log.Printf("failed to parse JSON: %v", err)
		return emptyApi, err
	}

	return &DeployAPI{
		Deployments: deployments,
	}, nil
}
