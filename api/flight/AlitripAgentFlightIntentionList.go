package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightIntentionList 意向单列表
// alitrip.agent.flight.intention.list
//
// 意向单列表
func AlitripAgentFlightIntentionList(clt *core.SDKClient, req *flight.AlitripAgentFlightIntentionListAPIRequest, session string) (*flight.AlitripAgentFlightIntentionListAPIResponse, error) {
	var resp flight.AlitripAgentFlightIntentionListAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
