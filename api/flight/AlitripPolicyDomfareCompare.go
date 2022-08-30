package flight

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyDomfareCompare 比价工具
// alitrip.policy.domfare.compare
//
// 比价工具
func AlitripPolicyDomfareCompare(clt *core.SDKClient, req *flight.AlitripPolicyDomfareCompareAPIRequest, session string) (*flight.AlitripPolicyDomfareCompareAPIResponse, error) {
	var resp flight.AlitripPolicyDomfareCompareAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
