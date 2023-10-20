package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripOrderGetAPIRequest 获取欢行统一订单模型 API请求
// alibaba.happytrip.order.get
//
// 通过订单id获取欢行统一订单模型数据
type AlibabaHappytripOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabaHappytripOrderGetRequest 初始化AlibabaHappytripOrderGetAPIRequest对象
func NewAlibabaHappytripOrderGetRequest() *AlibabaHappytripOrderGetAPIRequest {
	return &AlibabaHappytripOrderGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripOrderGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.order.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripOrderGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripOrderGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaHappytripOrderGetAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripOrderGetAPIRequest) GetOrderId() int64 {
	return r._orderId
}
