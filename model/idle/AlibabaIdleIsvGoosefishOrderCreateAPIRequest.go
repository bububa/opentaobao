package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvGoosefishOrderCreateAPIRequest 闲鱼三方安康容器订单创建 API请求
// alibaba.idle.isv.goosefish.order.create
//
// 闲鱼三方安康容器订单创建
type AlibabaIdleIsvGoosefishOrderCreateAPIRequest struct {
	model.Params
	// 创单请求参数
	_orderCreateRequest *OrderCreateRequest
}

// NewAlibabaIdleIsvGoosefishOrderCreateRequest 初始化AlibabaIdleIsvGoosefishOrderCreateAPIRequest对象
func NewAlibabaIdleIsvGoosefishOrderCreateRequest() *AlibabaIdleIsvGoosefishOrderCreateAPIRequest {
	return &AlibabaIdleIsvGoosefishOrderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleIsvGoosefishOrderCreateAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.isv.goosefish.order.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleIsvGoosefishOrderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleIsvGoosefishOrderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderCreateRequest is OrderCreateRequest Setter
// 创单请求参数
func (r *AlibabaIdleIsvGoosefishOrderCreateAPIRequest) SetOrderCreateRequest(_orderCreateRequest *OrderCreateRequest) error {
	r._orderCreateRequest = _orderCreateRequest
	r.Set("order_create_request", _orderCreateRequest)
	return nil
}

// GetOrderCreateRequest OrderCreateRequest Getter
func (r AlibabaIdleIsvGoosefishOrderCreateAPIRequest) GetOrderCreateRequest() *OrderCreateRequest {
	return r._orderCreateRequest
}
