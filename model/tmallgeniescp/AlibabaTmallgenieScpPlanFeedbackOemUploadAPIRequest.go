package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest
14-供应商反馈（OEM）同步接口 API请求
alibaba.tmallgenie.scp.plan.feedback.oem.upload

供应商反馈（OEM）同步接口 */
type AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
