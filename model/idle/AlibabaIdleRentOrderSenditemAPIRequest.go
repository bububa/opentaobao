package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaidlerentordersenditemAPIRequest 确认发货 API请求
// alibaba.idle.rent.order.senditem
//
// 确认发货
type AlibabaidlerentordersenditemAPIRequest struct {
	model.Params
	// 物流信息
	_logisticsList []LogisticsDto
	// 订单id
	_orderId int64
}

// NewAlibabaidlerentordersenditemRequest 初始化AlibabaidlerentordersenditemAPIRequest对象
func NewAlibabaidlerentordersenditemRequest() *AlibabaidlerentordersenditemAPIRequest {
	return &AlibabaidlerentordersenditemAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaidlerentordersenditemAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.senditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaidlerentordersenditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaidlerentordersenditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsList is LogisticsList Setter
// 物流信息
func (r *AlibabaidlerentordersenditemAPIRequest) SetLogisticsList(_logisticsList []LogisticsDto) error {
	r._logisticsList = _logisticsList
	r.Set("logistics_list", _logisticsList)
	return nil
}

// GetLogisticsList LogisticsList Getter
func (r AlibabaidlerentordersenditemAPIRequest) GetLogisticsList() []LogisticsDto {
	return r._logisticsList
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaidlerentordersenditemAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaidlerentordersenditemAPIRequest) GetOrderId() int64 {
	return r._orderId
}
