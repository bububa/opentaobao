package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripAgentFlightIntentionConfirm 意向单确认
// alitrip.agent.flight.intention.confirm
//
// 意向单确认
func AlitripAgentFlightIntentionConfirm(clt *core.SDKClient, req *flight.AlitripAgentFlightIntentionConfirmAPIRequest, session string) (*flight.AlitripAgentFlightIntentionConfirmAPIResponse, error) {
	var resp flight.AlitripAgentFlightIntentionConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
