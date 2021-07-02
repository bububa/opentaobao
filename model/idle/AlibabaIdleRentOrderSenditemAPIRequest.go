package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderSenditemAPIRequest 确认发货 API请求
// alibaba.idle.rent.order.senditem
//
// 确认发货
type AlibabaIdleRentOrderSenditemAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logisticsList []LogisticsDto
}

// NewAlibabaIdleRentOrderSenditemRequest 初始化AlibabaIdleRentOrderSenditemAPIRequest对象
func NewAlibabaIdleRentOrderSenditemRequest() *AlibabaIdleRentOrderSenditemAPIRequest {
	return &AlibabaIdleRentOrderSenditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.senditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderSenditemAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is LogisticsList Setter
// 物流信息
func (r *AlibabaIdleRentOrderSenditemAPIRequest) SetLogisticsList(_logisticsList []LogisticsDto) error {
	r._logisticsList = _logisticsList
	r.Set("logistics_list", _logisticsList)
	return nil
}

// Get LogisticsList Getter
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetLogisticsList() []LogisticsDto {
	return r._logisticsList
}
