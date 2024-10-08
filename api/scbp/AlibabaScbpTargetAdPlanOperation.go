package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanOperation 定向推广-计划开启/暂停/删除
// alibaba.scbp.target.ad.plan.operation
//
// 定向推广-计划开启/暂停/删除
func AlibabaScbpTargetAdPlanOperation(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanOperationAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanOperationAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
