package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanCrowdIdGet 定向推广-人群标签ID获取(店铺老客、优选人群)
// alibaba.scbp.target.ad.plan.crowd.id.get
//
// 定向推广-人群标签ID获取(店铺老客、优选人群)
func AlibabaScbpTargetAdPlanCrowdIdGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanCrowdIdGetAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
