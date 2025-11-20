package deploy

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Tata-Matata/fake-aria-api/util"
)

type DeployAPI struct {
	Deployments []interface{} `json:"deployments"`
}

// generates deployments off of original json deployment by randomizing values
// the number of deployments is picked randomly from the range min-max, max excluded
func (api *DeployAPI) Randomize() ([]interface{}, error) {

	//how many deploys to generate
	total := util.RandomFromRange(MIN_RANDOM_DEPLOYS, MAX_RANDOM_DEPLOYS)

	//generate list of random deploys
	deploys := make([]interface{}, 0, total) // length 0, capacity = how many random deploys
	for i := 0; i < total; i++ {
		deploy, err := api.createRandomDeploy(i)
		if err != nil {
			log.Printf("Failed to create deploy")
		} else {
			deploys = append(deploys, deploy)
		}
	}

	return deploys, nil
}

// creates a single random deployment by picking a random deployment from the json
func (api *DeployAPI) createRandomDeploy(counter int) (interface{}, error) {
	//pick random deployment from json
	idx := util.RandomFromRange(0, len(api.Deployments))

	//create deep copy of deployment
	deploy, err := util.DeepCopy(api.Deployments[idx])
	if err != nil {
		return nil, fmt.Errorf("failed to create deep copy of deployment: %v", err)
	}

	//set random project
	projectId, err := util.RandomFromMap(PROJECTS)
	if err != nil {
		return nil, fmt.Errorf("failed to pick random project: %v", err)
	}
	deploy, err = util.SetStringField(deploy, PROJID, projectId)
	if err != nil {
		return nil, fmt.Errorf("failed to set projectId: %v", err)
	}

	//set random deployment status
	status := util.RandomFromList(STATUSES)
	deploy, err = util.SetStringField(deploy, STATUS, status)
	if err != nil {
		return nil, fmt.Errorf("failed to set status: %v", err)
	}

	//set random ID
	deploy, err = util.SetStringField(deploy, ID, util.RandomUuid())
	if err != nil {
		return nil, fmt.Errorf("failed to set id: %v", err)
	}

	//set name
	deploy, err = util.SetStringField(deploy, NAME, NAME_PREFIX+fmt.Sprintf("%d", counter))
	if err != nil {
		return nil, fmt.Errorf("failed to set name: %v", err)
	}

	return deploy, nil
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
	message := fmt.Sprintf("deployment with id %v not found", id)
	log.Print(message)
	return nil, fmt.Errorf("%s", message)
}

func NewDeployAPI(jsonPath string) (*DeployAPI, error) {

	emptyApi := &DeployAPI{
		Deployments: []interface{}{},
	}

	data, err := os.ReadFile(jsonPath)
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
