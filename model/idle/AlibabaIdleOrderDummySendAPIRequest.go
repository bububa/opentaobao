package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleOrderDummySendAPIRequest 闲鱼无需物流发货 API请求
// alibaba.idle.order.dummy.send
//
// 适用于电子卡券等虚拟商品不需要物流的商品发货。
type AlibabaIdleOrderDummySendAPIRequest struct {
	model.Params
	// 请求
	_paramOrderDummySendRequest *OrderDummySendRequest
}

// NewAlibabaIdleOrderDummySendRequest 初始化AlibabaIdleOrderDummySendAPIRequest对象
func NewAlibabaIdleOrderDummySendRequest() *AlibabaIdleOrderDummySendAPIRequest {
	return &AlibabaIdleOrderDummySendAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleOrderDummySendAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.order.dummy.send"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleOrderDummySendAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleOrderDummySendAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamOrderDummySendRequest is ParamOrderDummySendRequest Setter
// 请求
func (r *AlibabaIdleOrderDummySendAPIRequest) SetParamOrderDummySendRequest(_paramOrderDummySendRequest *OrderDummySendRequest) error {
	r._paramOrderDummySendRequest = _paramOrderDummySendRequest
	r.Set("param_order_dummy_send_request", _paramOrderDummySendRequest)
	return nil
}

// GetParamOrderDummySendRequest ParamOrderDummySendRequest Getter
func (r AlibabaIdleOrderDummySendAPIRequest) GetParamOrderDummySendRequest() *OrderDummySendRequest {
	return r._paramOrderDummySendRequest
}
