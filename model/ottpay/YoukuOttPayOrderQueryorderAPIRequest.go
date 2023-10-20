package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuOttPayOrderQueryorderAPIRequest 查询订单 API请求
// youku.ott.pay.order.queryorder
//
// 通过订单号查询订单信息
type YoukuOttPayOrderQueryorderAPIRequest struct {
	model.Params
	// 订单号
	_orderNo string
}

// NewYoukuOttPayOrderQueryorderRequest 初始化YoukuOttPayOrderQueryorderAPIRequest对象
func NewYoukuOttPayOrderQueryorderRequest() *YoukuOttPayOrderQueryorderAPIRequest {
	return &YoukuOttPayOrderQueryorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuOttPayOrderQueryorderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuOttPayOrderQueryorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuOttPayOrderQueryorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单号
func (r *YoukuOttPayOrderQueryorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuOttPayOrderQueryorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}
