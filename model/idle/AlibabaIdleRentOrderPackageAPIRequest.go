package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderPackageAPIRequest 确认揽收商品 API请求
// alibaba.idle.rent.order.package
//
// 确认揽收
type AlibabaIdleRentOrderPackageAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logistics *LogisticsDto
}

// NewAlibabaIdleRentOrderPackageRequest 初始化AlibabaIdleRentOrderPackageAPIRequest对象
func NewAlibabaIdleRentOrderPackageRequest() *AlibabaIdleRentOrderPackageAPIRequest {
	return &AlibabaIdleRentOrderPackageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderPackageAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.package"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderPackageAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderPackageAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaIdleRentOrderPackageAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// Set is Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderPackageAPIRequest) SetLogistics(_logistics *LogisticsDto) error {
	r._logistics = _logistics
	r.Set("logistics", _logistics)
	return nil
}

// Get Logistics Getter
func (r AlibabaIdleRentOrderPackageAPIRequest) GetLogistics() *LogisticsDto {
	return r._logistics
}
