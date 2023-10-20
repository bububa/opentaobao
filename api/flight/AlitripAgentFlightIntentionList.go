package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// Alitripagentflightintentionlist 意向单列表
// alitrip.agent.flight.intention.list
//
// 意向单列表
func Alitripagentflightintentionlist(clt *core.SDKClient, req *flight.AlitripagentflightintentionlistAPIRequest, session string) (*flight.AlitripagentflightintentionlistAPIResponse, error) {
	var resp flight.AlitripagentflightintentionlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
