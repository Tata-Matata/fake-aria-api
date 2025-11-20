package storage

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Tata-Matata/fake-aria-api/util"
)

type DataStore struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Capacity uint64 `json:"capacity"`
}
type StorageAPI struct {
	DataStores []DataStore `json:"dataStores"`
}

// generates data stores randomizing capacity values
func (api *StorageAPI) Randomize() ([]DataStore, error) {
	for i, dstore := range api.DataStores {
		//randomize capacity
		capacity := util.RandomFromRange(MIN_RANDOM_CAPACITY, MAX_RANDOM_CAPACITY)
		api.DataStores[i].Capacity = uint64(capacity)
		log.Printf("data store %v capacity set to %v", dstore.ID, util.BytesToTiB(uint64(capacity)))
	}
	return api.DataStores, nil
}

// retrieves data store by ID
func (api *StorageAPI) GetByID(id string) (DataStore, error) {

	for _, dstore := range api.DataStores {

		if dstore.ID == id {
			return dstore, nil
		}
	}
	message := fmt.Sprintf("data store with id %v not found", id)
	log.Print(message)
	return DataStore{}, fmt.Errorf("%s", message)
}

// initializes StorageAPI from JSON file
func NewStorageAPI(jsonPath string) (*StorageAPI, error) {

	//empty API for error cases
	dataStores := []DataStore{}
	emptyApi := &StorageAPI{
		DataStores: dataStores,
	}
	//read JSON file
	data, err := os.ReadFile(jsonPath)
	if err != nil {
		log.Printf("failed to read JSON file: %v", err)
		return emptyApi, err
	}

	//parse JSON
	if err := json.Unmarshal(data, &dataStores); err != nil {
		log.Printf("failed to parse JSON: %v", err)
		return emptyApi, err
	}

	return &StorageAPI{
		DataStores: dataStores,
	}, nil
}
