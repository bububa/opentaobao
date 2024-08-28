package flight

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/flight"
)

// AlitripPolicyDomfareCompare 比价工具
// alitrip.policy.domfare.compare
//
// 比价工具
func AlitripPolicyDomfareCompare(ctx context.Context, clt *core.SDKClient, req *flight.AlitripPolicyDomfareCompareAPIRequest, resp *flight.AlitripPolicyDomfareCompareAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
