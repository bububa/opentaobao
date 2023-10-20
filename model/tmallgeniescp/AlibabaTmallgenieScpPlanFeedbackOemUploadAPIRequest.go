package tmallgeniescp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest 14-供应商反馈（OEM）同步接口 API请求
// alibaba.tmallgenie.scp.plan.feedback.oem.upload
//
// 供应商反馈（OEM）同步接口
type AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest struct {
	model.Params
	// 扩展参数
	_requestExtendJson string
}

// NewAlibabatmallgeniescpplanfeedbackoemuploadRequest 初始化AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest对象
func NewAlibabatmallgeniescpplanfeedbackoemuploadRequest() *AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest {
	return &AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.tmallgenie.scp.plan.feedback.oem.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestExtendJson is RequestExtendJson Setter
// 扩展参数
func (r *AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest) SetRequestExtendJson(_requestExtendJson string) error {
	r._requestExtendJson = _requestExtendJson
	r.Set("request_extend_json", _requestExtendJson)
	return nil
}

// GetRequestExtendJson RequestExtendJson Getter
func (r AlibabatmallgeniescpplanfeedbackoemuploadAPIRequest) GetRequestExtendJson() string {
	return r._requestExtendJson
}
