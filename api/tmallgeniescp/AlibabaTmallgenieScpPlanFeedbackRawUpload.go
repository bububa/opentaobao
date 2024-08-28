package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanFeedbackRawUpload 15-供应商反馈（原料）同步接口
// alibaba.tmallgenie.scp.plan.feedback.raw.upload
//
// 供应商反馈（原料）同步接口
func AlibabaTmallgenieScpPlanFeedbackRawUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackRawUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
