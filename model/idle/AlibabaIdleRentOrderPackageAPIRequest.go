package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentorderpackageAPIRequest 确认揽收商品 API请求
// alibaba.idle.rent.order.package
//
// 确认揽收
type AlibabaidlerentorderpackageAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logistics *LogisticsDto
}

// NewAlibabaidlerentorderpackageRequest 初始化AlibabaidlerentorderpackageAPIRequest对象
func NewAlibabaidlerentorderpackageRequest() *AlibabaidlerentorderpackageAPIRequest {
	return &AlibabaidlerentorderpackageAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentorderpackageAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.package"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentorderpackageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentorderpackageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaidlerentorderpackageAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaidlerentorderpackageAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetLogistics is Logistics Setter
// 物流信息
func (r *AlibabaidlerentorderpackageAPIRequest) SetLogistics(_logistics *LogisticsDto) error {
	r._logistics = _logistics
	r.Set("logistics", _logistics)
	return nil
}

// GetLogistics Logistics Getter
func (r AlibabaidlerentorderpackageAPIRequest) GetLogistics() *LogisticsDto {
	return r._logistics
}
