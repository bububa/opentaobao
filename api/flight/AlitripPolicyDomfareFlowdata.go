package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyDomfareFlowdata 比价工具流量详情
// alitrip.policy.domfare.flowdata
//
// 比价工具流量详情
func AlitripPolicyDomfareFlowdata(clt *core.SDKClient, req *flight.AlitripPolicyDomfareFlowdataAPIRequest, resp *flight.AlitripPolicyDomfareFlowdataAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
