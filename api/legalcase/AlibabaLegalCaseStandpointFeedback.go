package legalcase

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/legalcase"
)

// AlibabaLegalCaseStandpointFeedback 新增或更新 反馈口径(采纳口径/不采纳口径)
// alibaba.legal.case.standpoint.feedback
//
// 新增或更新 反馈口径(采纳口径/不采纳口径)
func AlibabaLegalCaseStandpointFeedback(ctx context.Context, clt *core.SDKClient, req *legalcase.AlibabaLegalCaseStandpointFeedbackAPIRequest, resp *legalcase.AlibabaLegalCaseStandpointFeedbackAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
