package deploy

const DEPLOY_JSON string = "data/deployments.json"

var PROJECTS = map[string]string{
	"Project1": "c0917929-ff03-46f8-853b-066232a9343a",
	"Project2": "d0017929-ff03-46f8-353b-066232a93bcd",
	"Proejct3": "e1118930-ff03-46f8-353b-054232a13bxc",
}

var STATUSES = []string{
	"CREATE_PENDING",
	"CREATE_SUCCESSFUL",
	"CREATE_FAILED",
	"DELETE_FAILED",
}

// for randomizing deployments
const MIN_RANDOM_DEPLOYS int = 10
const MAX_RANDOM_DEPLOYS int = 300
const NAME_PREFIX string = "Deploy-"

// deploy object fields
const PROJID string = "projectId"
const STATUS string = "status"
const ID string = "id"
const NAME string = "name"
