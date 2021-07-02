package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscOrderOrderUploadAPIRequest 订单回流 API请求
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
type AlibabaAlscOrderOrderUploadAPIRequest struct {
	model.Params
	// 订单回流参数
	_paramBackflowRequest *BackflowRequest
}

// NewAlibabaAlscOrderOrderUploadRequest 初始化AlibabaAlscOrderOrderUploadAPIRequest对象
func NewAlibabaAlscOrderOrderUploadRequest() *AlibabaAlscOrderOrderUploadAPIRequest {
	return &AlibabaAlscOrderOrderUploadAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetApiMethodName() string {
	return "alibaba.alsc.order.order.upload"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetParamBackflowRequest is ParamBackflowRequest Setter
// 订单回流参数
func (r *AlibabaAlscOrderOrderUploadAPIRequest) SetParamBackflowRequest(_paramBackflowRequest *BackflowRequest) error {
	r._paramBackflowRequest = _paramBackflowRequest
	r.Set("param_backflow_request", _paramBackflowRequest)
	return nil
}

// GetParamBackflowRequest ParamBackflowRequest Getter
func (r AlibabaAlscOrderOrderUploadAPIRequest) GetParamBackflowRequest() *BackflowRequest {
	return r._paramBackflowRequest
}
