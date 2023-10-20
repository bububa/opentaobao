package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalscorderorderuploadAPIRequest 订单回流 API请求
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
type AlibabaalscorderorderuploadAPIRequest struct {
	model.Params
	// 订单回流参数
	_paramBackflowRequest *BackflowRequest
}

// NewAlibabaalscorderorderuploadRequest 初始化AlibabaalscorderorderuploadAPIRequest对象
func NewAlibabaalscorderorderuploadRequest() *AlibabaalscorderorderuploadAPIRequest {
	return &AlibabaalscorderorderuploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalscorderorderuploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.order.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalscorderorderuploadAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalscorderorderuploadAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamBackflowRequest is ParamBackflowRequest Setter
// 订单回流参数
func (r *AlibabaalscorderorderuploadAPIRequest) SetParamBackflowRequest(_paramBackflowRequest *BackflowRequest) error {
	r._paramBackflowRequest = _paramBackflowRequest
	r.Set("param_backflow_request", _paramBackflowRequest)
	return nil
}

// GetParamBackflowRequest ParamBackflowRequest Getter
func (r AlibabaalscorderorderuploadAPIRequest) GetParamBackflowRequest() *BackflowRequest {
	return r._paramBackflowRequest
}
