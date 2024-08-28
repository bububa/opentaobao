package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanList 定向推广-查询定向推广计划列表并返回计划基础信息
// alibaba.scbp.target.ad.plan.list
//
// 定向推广-查询定向推广计划列表并返回计划基础信息
func AlibabaScbpTargetAdPlanList(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanListAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
