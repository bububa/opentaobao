package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyDomfareFlowdata 比价工具流量详情
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
func AlitripPolicyDomfareFlowdata(clt *core.SDKClient, req *flight.AlitripPolicyDomfareFlowdataAPIRequest, session string) (*flight.AlitripPolicyDomfareFlowdataAPIResponse, error) {
	var resp flight.AlitripPolicyDomfareFlowdataAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
