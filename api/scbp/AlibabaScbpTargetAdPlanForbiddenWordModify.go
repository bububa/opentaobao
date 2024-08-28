package scbp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scbp"
)

// AlibabaScbpTargetAdPlanForbiddenWordModify 定向推广-新增或删除屏蔽词
// alibaba.scbp.target.ad.plan.forbidden.word.modify
//
// 定向推广-新增或删除屏蔽词
func AlibabaScbpTargetAdPlanForbiddenWordModify(ctx context.Context, clt *core.SDKClient, req *scbp.AlibabaScbpTargetAdPlanForbiddenWordModifyAPIRequest, resp *scbp.AlibabaScbpTargetAdPlanForbiddenWordModifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
