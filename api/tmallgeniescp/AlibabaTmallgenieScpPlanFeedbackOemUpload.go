package tmallgeniescp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanFeedbackOemUpload 14-供应商反馈（OEM）同步接口
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
func AlibabaTmallgenieScpPlanFeedbackOemUpload(ctx context.Context, clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest, resp *tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
