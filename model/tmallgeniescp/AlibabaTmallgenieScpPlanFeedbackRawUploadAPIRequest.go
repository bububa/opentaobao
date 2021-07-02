package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest 15-供应商反馈（原料）同步接口 API请求
// alibaba.tmallgenie.scp.plan.feedback.raw.upload
//
// 供应商反馈（原料）同步接口
type AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest 初始化AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest对象
func NewAlibabaTmallgenieScpPlanFeedbackRawUploadRequest() *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest {
	return &AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.feedback.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestExtendJson Setter
// 扩展参数
func (r *AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// Get RequestExtendJson Getter
func (r AlibabaTmallgenieScpPlanFeedbackRawUploadAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
