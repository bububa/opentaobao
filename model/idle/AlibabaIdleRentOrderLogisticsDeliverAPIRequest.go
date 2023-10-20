package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentorderlogisticsdeliverAPIRequest 创建揽收物流 API请求
// alibaba.idle.rent.order.logistics.deliver
//
// 创建揽收物流
// 商家去物流公司创建物流订单
type AlibabaidlerentorderlogisticsdeliverAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logistics *LogisticsDto
}

// NewAlibabaidlerentorderlogisticsdeliverRequest 初始化AlibabaidlerentorderlogisticsdeliverAPIRequest对象
func NewAlibabaidlerentorderlogisticsdeliverRequest() *AlibabaidlerentorderlogisticsdeliverAPIRequest {
	return &AlibabaidlerentorderlogisticsdeliverAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentorderlogisticsdeliverAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.logistics.deliver"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentorderlogisticsdeliverAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentorderlogisticsdeliverAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaidlerentorderlogisticsdeliverAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaidlerentorderlogisticsdeliverAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetLogistics is Logistics Setter
// 物流信息
func (r *AlibabaidlerentorderlogisticsdeliverAPIRequest) SetLogistics(_logistics *LogisticsDto) error {
	r._logistics = _logistics
	r.Set("logistics", _logistics)
	return nil
}

// GetLogistics Logistics Getter
func (r AlibabaidlerentorderlogisticsdeliverAPIRequest) GetLogistics() *LogisticsDto {
	return r._logistics
}
