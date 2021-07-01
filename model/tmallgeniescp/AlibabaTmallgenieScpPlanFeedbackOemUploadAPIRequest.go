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

// NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest 初始化AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackOemUploadRequest() *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.feedback.oem.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// Get RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackOemUploadAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
