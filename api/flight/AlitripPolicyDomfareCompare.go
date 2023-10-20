package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyDomfareCompare 比价工具
// alitrip.policy.domfare.compare
//
// 比价工具
func AlitripPolicyDomfareCompare(clt *core.SDKClient, req *flight.AlitripPolicyDomfareCompareAPIRequest, resp *flight.AlitripPolicyDomfareCompareAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
