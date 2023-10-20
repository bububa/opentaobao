package ottpay

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YoukuottpayorderqueryorderAPIRequest 查询订单 API请求
// youku.ott.pay.order.queryorder
//
// 通过订单号查询订单信息
type YoukuottpayorderqueryorderAPIRequest struct {
	model.Params
	// 订单号
	_orderNo string
}

// NewYoukuottpayorderqueryorderRequest 初始化YoukuottpayorderqueryorderAPIRequest对象
func NewYoukuottpayorderqueryorderRequest() *YoukuottpayorderqueryorderAPIRequest {
	return &YoukuottpayorderqueryorderAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YoukuottpayorderqueryorderAPIRequest) GetApiMethodName() string {
	return "youku.ott.pay.order.queryorder"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YoukuottpayorderqueryorderAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YoukuottpayorderqueryorderAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderNo is OrderNo Setter
// 订单号
func (r *YoukuottpayorderqueryorderAPIRequest) SetOrderNo(_orderNo string) error {
	r._orderNo = _orderNo
	r.Set("order_no", _orderNo)
	return nil
}

// GetOrderNo OrderNo Getter
func (r YoukuottpayorderqueryorderAPIRequest) GetOrderNo() string {
	return r._orderNo
}
