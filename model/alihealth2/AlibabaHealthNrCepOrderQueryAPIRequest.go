package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahealthnrceporderqueryAPIRequest 订单详情查询接口 API请求
// alibaba.health.nr.cep.order.query
//
// 订单详情查询接口
type AlibabahealthnrceporderqueryAPIRequest struct {
	model.Params
	// 订单号
	_orderId int64
}

// NewAlibabahealthnrceporderqueryRequest 初始化AlibabahealthnrceporderqueryAPIRequest对象
func NewAlibabahealthnrceporderqueryRequest() *AlibabahealthnrceporderqueryAPIRequest {
	return &AlibabahealthnrceporderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahealthnrceporderqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.health.nr.cep.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahealthnrceporderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahealthnrceporderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单号
func (r *AlibabahealthnrceporderqueryAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahealthnrceporderqueryAPIRequest) GetOrderId() int64 {
	return r._orderId
}
