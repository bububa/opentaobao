package aliexpresssumaitong

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressTradeOrderOpenCheckAPIRequest Aliexpress开放平台下单前置检查 API请求
// aliexpress.trade.order.open.check
//
// Aliexpress开放平台下单前通过下单入参获取token
type AliexpressTradeOrderOpenCheckAPIRequest struct {
	model.Params
	// 预下单入参
	_paramPreCreateOrderRequest *PreCreateOrderRequest
	// 客户端信息
	_paramClientInfo *ClientInfo
}

// NewAliexpressTradeOrderOpenCheckRequest 初始化AliexpressTradeOrderOpenCheckAPIRequest对象
func NewAliexpressTradeOrderOpenCheckRequest() *AliexpressTradeOrderOpenCheckAPIRequest {
	return &AliexpressTradeOrderOpenCheckAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressTradeOrderOpenCheckAPIRequest) Reset() {
	r._paramPreCreateOrderRequest = nil
	r._paramClientInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetApiMethodName() string {
	return "aliexpress.trade.order.open.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamPreCreateOrderRequest is ParamPreCreateOrderRequest Setter
// 预下单入参
func (r *AliexpressTradeOrderOpenCheckAPIRequest) SetParamPreCreateOrderRequest(_paramPreCreateOrderRequest *PreCreateOrderRequest) error {
	r._paramPreCreateOrderRequest = _paramPreCreateOrderRequest
	r.Set("param_pre_create_order_request", _paramPreCreateOrderRequest)
	return nil
}

// GetParamPreCreateOrderRequest ParamPreCreateOrderRequest Getter
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetParamPreCreateOrderRequest() *PreCreateOrderRequest {
	return r._paramPreCreateOrderRequest
}

// SetParamClientInfo is ParamClientInfo Setter
// 客户端信息
func (r *AliexpressTradeOrderOpenCheckAPIRequest) SetParamClientInfo(_paramClientInfo *ClientInfo) error {
	r._paramClientInfo = _paramClientInfo
	r.Set("param_client_info", _paramClientInfo)
	return nil
}

// GetParamClientInfo ParamClientInfo Getter
func (r AliexpressTradeOrderOpenCheckAPIRequest) GetParamClientInfo() *ClientInfo {
	return r._paramClientInfo
}

var poolAliexpressTradeOrderOpenCheckAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressTradeOrderOpenCheckRequest()
	},
}

// GetAliexpressTradeOrderOpenCheckRequest 从 sync.Pool 获取 AliexpressTradeOrderOpenCheckAPIRequest
func GetAliexpressTradeOrderOpenCheckAPIRequest() *AliexpressTradeOrderOpenCheckAPIRequest {
	return poolAliexpressTradeOrderOpenCheckAPIRequest.Get().(*AliexpressTradeOrderOpenCheckAPIRequest)
}

// ReleaseAliexpressTradeOrderOpenCheckAPIRequest 将 AliexpressTradeOrderOpenCheckAPIRequest 放入 sync.Pool
func ReleaseAliexpressTradeOrderOpenCheckAPIRequest(v *AliexpressTradeOrderOpenCheckAPIRequest) {
	v.Reset()
	poolAliexpressTradeOrderOpenCheckAPIRequest.Put(v)
}
