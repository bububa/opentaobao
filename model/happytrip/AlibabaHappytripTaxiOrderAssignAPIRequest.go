package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiorderassignAPIRequest 订单指派 API请求
// alibaba.happytrip.taxi.order.assign
//
// 通知供应商订单指派成功
type AlibabahappytriptaxiorderassignAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
}

// NewAlibabahappytriptaxiorderassignRequest 初始化AlibabahappytriptaxiorderassignAPIRequest对象
func NewAlibabahappytriptaxiorderassignRequest() *AlibabahappytriptaxiorderassignAPIRequest {
	return &AlibabahappytriptaxiorderassignAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiorderassignAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.assign"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiorderassignAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiorderassignAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商订单号
func (r *AlibabahappytriptaxiorderassignAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiorderassignAPIRequest) GetOrderId() string {
	return r._orderId
}
