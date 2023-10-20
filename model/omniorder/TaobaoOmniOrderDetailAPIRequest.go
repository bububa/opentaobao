package omniorder

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoomniorderdetailAPIRequest 全渠道订单详情 API请求
// taobao.omni.order.detail
//
// 全渠道订单详情
type TaobaoomniorderdetailAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewTaobaoomniorderdetailRequest 初始化TaobaoomniorderdetailAPIRequest对象
func NewTaobaoomniorderdetailRequest() *TaobaoomniorderdetailAPIRequest {
	return &TaobaoomniorderdetailAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoomniorderdetailAPIRequest) GetApiMethodName() string {
	return "taobao.omni.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoomniorderdetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoomniorderdetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *TaobaoomniorderdetailAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r TaobaoomniorderdetailAPIRequest) GetOrderId() int64 {
	return r._orderId
}
