package deploy

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomize(t *testing.T) {

	curWorkDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	api, err := NewDeployAPI(curWorkDir + "/../" + DEPLOY_JSON)

	assert.NoError(t, err)

	deploys, err := api.Randomize()
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(deploys), MIN_RANDOM_DEPLOYS)
	assert.Less(t, len(deploys), MAX_RANDOM_DEPLOYS)

	// Check for unique names and IDs
	nameSet := make(map[string]bool)
	idSet := make(map[string]bool)
	projectIDSet := make(map[string]bool)
	statusSet := make(map[string]bool)

	for _, deploy := range deploys {
		deployMap, ok := deploy.(map[string]interface{})
		assert.True(t, ok, "deployment should be a map")

		// Check uniqueness of names
		name, nameOk := deployMap[NAME].(string)
		assert.True(t, nameOk, "deployment should have a name")
		assert.False(t, nameSet[name], "name should be unique")
		nameSet[name] = true

		// Check uniqueness of IDs
		id, idOk := deployMap[ID].(string)
		assert.True(t, idOk, "deployment should have an ID")
		assert.False(t, idSet[id], "ID should be unique")
		idSet[id] = true

		// Collect project IDs
		projectID, projectIDOk := deployMap[PROJID].(string)
		if projectIDOk {
			projectIDSet[projectID] = true
			// Check if project ID is valid
			assert.True(t, func() bool { _, ok := PROJECTS[projectID]; return ok }())
		}

		// Check if status is valid
		status, statusOk := deployMap[STATUS].(string)
		if statusOk {
			statusSet[status] = true
			assert.Contains(t, STATUSES, status, "status should be one of the predefined STATUSES")
		}
	}

	// Ensure there are at least some different statuses and project IDs
	assert.GreaterOrEqual(t, len(statusSet), 2, "there should be at least two different statuses")
	assert.GreaterOrEqual(t, len(projectIDSet), 2, "there should be at least two different project IDs")

}
