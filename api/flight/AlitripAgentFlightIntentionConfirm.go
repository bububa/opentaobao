package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightintentionconfirm 意向单确认
// alitrip.agent.flight.intention.confirm
//
// 意向单确认
func Alitripagentflightintentionconfirm(clt *core.SDKClient, req *flight.AlitripagentflightintentionconfirmAPIRequest, session string) (*flight.AlitripagentflightintentionconfirmAPIResponse, error) {
	var resp flight.AlitripagentflightintentionconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
