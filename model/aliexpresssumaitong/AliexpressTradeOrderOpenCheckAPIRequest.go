package aliexpresssumaitong

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresstradeorderopencheckAPIRequest Aliexpress开放平台下单前置检查 API请求
// aliexpress.trade.order.open.check
//
// Aliexpress开放平台下单前通过下单入参获取token
type AliexpresstradeorderopencheckAPIRequest struct {
	model.Params
	// 预下单入参
	_paramPreCreateOrderRequest *PreCreateOrderRequest
	// 客户端信息
	_paramClientInfo *ClientInfo
}

// NewAliexpresstradeorderopencheckRequest 初始化AliexpresstradeorderopencheckAPIRequest对象
func NewAliexpresstradeorderopencheckRequest() *AliexpresstradeorderopencheckAPIRequest {
	return &AliexpresstradeorderopencheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresstradeorderopencheckAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.order.open.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresstradeorderopencheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresstradeorderopencheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPreCreateOrderRequest is ParamPreCreateOrderRequest Setter
// 预下单入参
func (r *AliexpresstradeorderopencheckAPIRequest) SetParamPreCreateOrderRequest(_paramPreCreateOrderRequest *PreCreateOrderRequest) error {
	r._paramPreCreateOrderRequest = _paramPreCreateOrderRequest
	r.Set("param_pre_create_order_request", _paramPreCreateOrderRequest)
	return nil
}

// GetParamPreCreateOrderRequest ParamPreCreateOrderRequest Getter
func (r AliexpresstradeorderopencheckAPIRequest) GetParamPreCreateOrderRequest() *PreCreateOrderRequest {
	return r._paramPreCreateOrderRequest
}

// SetParamClientInfo is ParamClientInfo Setter
// 客户端信息
func (r *AliexpresstradeorderopencheckAPIRequest) SetParamClientInfo(_paramClientInfo *ClientInfo) error {
	r._paramClientInfo = _paramClientInfo
	r.Set("param_client_info", _paramClientInfo)
	return nil
}

// GetParamClientInfo ParamClientInfo Getter
func (r AliexpresstradeorderopencheckAPIRequest) GetParamClientInfo() *ClientInfo {
	return r._paramClientInfo
}
