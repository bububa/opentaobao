package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// AlibabaTmallgenieScpPlanFeedbackOemUpload 14-供应商反馈（OEM）同步接口
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
func AlibabaTmallgenieScpPlanFeedbackOemUpload(clt *core.SDKClient, req *tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest, session string) (*tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabaTmallgenieScpPlanFeedbackOemUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
