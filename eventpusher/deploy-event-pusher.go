package eventpusher

import (
	"time"

	"github.com/Tata-Matata/fake-aria-api/util"
)

type DeployEvent struct {
	Status    string `json:"status"`
	ErrorType string `json:"errorType"`
}

type DeployEventPusher struct{}

func (p DeployEventPusher) Name() string {
	return "StatusEvent"
}

func (p DeployEventPusher) Interval() time.Duration {
	return DEPLOY_INTERVAL_SECONDS * time.Second
}

func (p DeployEventPusher) Endpoint() string {
	return "/event"
}

func (p DeployEventPusher) GenerateEvent() (any, error) {

	if time.Now().UnixNano()%4 == 0 {
		errType := util.RandomFromList(ERR_TYPES)
		return DeployEvent{Status: "failure", ErrorType: errType}, nil
	}
	return DeployEvent{Status: "success"}, nil
}
