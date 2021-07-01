package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest
15-供应商反馈（原料）同步接口 API请求
alibaba.tmallgenie.scp.plan.feedback.raw.upload

供应商反馈（原料）同步接口 */
type AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// New
