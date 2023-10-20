package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest 15-供应商反馈（原料）同步接口 API请求
// alibaba.tmallgenie.scp.plan.feedback.raw.upload
//
// 供应商反馈（原料）同步接口
type AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplanfeedbackrawuploadRequest 初始化AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest对象
func NewAlibabatmallgeniescpplanfeedbackrawuploadRequest() *AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest {
	return &AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.feedback.raw.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplanfeedbackrawuploadAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
