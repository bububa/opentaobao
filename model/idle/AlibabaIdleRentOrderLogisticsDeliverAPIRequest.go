package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderLogisticsDeliverAPIRequest 创建揽收物流 API请求
// alibaba.idle.rent.order.logistics.deliver
//
// 创建揽收物流
// 商家去物流公司创建物流订单
type AlibabaIdleRentOrderLogisticsDeliverAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logistics *LogisticsDto
}

// NewAlibabaIdleRentOrderLogisticsDeliverRequest 初始化AlibabaIdleRentOrderLogisticsDeliverAPIRequest对象
func NewAlibabaIdleRentOrderLogisticsDeliverRequest() *AlibabaIdleRentOrderLogisticsDeliverAPIRequest {
	return &AlibabaIdleRentOrderLogisticsDeliverAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.logistics.deliver"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderLogisticsDeliverAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderLogisticsDeliverAPIRequest) SetLogistics(_logistics *LogisticsDto) error {
	r._logistics = _logistics
	r.Set("logistics", _logistics)
	return nil
}

// Get Logistics Getter
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetLogistics() *LogisticsDto {
	return r._logistics
}
