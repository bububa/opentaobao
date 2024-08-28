package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanTagGet 定向推广-获取计划的定向溢价数据
// alibaba.scbp.target.ad.plan.tag.get
//
// 定向推广-获取计划的定向溢价数据
func AlibabaScbpTargetAdPlanTagGet(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanTagGetAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanTagGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
