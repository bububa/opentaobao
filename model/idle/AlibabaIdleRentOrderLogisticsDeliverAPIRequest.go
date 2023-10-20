package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentOrderLogisticsDeliverAPIRequest) Reset() {
	r._orderId = 0
	r._logistics = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.logistics.deliver"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderLogisticsDeliverAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetLogistics is Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderLogisticsDeliverAPIRequest) SetLogistics(_logistics *LogisticsDto) error {
	r._logistics = _logistics
	r.Set("logistics", _logistics)
	return nil
}

// GetLogistics Logistics Getter
func (r AlibabaIdleRentOrderLogisticsDeliverAPIRequest) GetLogistics() *LogisticsDto {
	return r._logistics
}

var poolAlibabaIdleRentOrderLogisticsDeliverAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentOrderLogisticsDeliverRequest()
	},
}

// GetAlibabaIdleRentOrderLogisticsDeliverRequest 从 sync.Pool 获取 AlibabaIdleRentOrderLogisticsDeliverAPIRequest
func GetAlibabaIdleRentOrderLogisticsDeliverAPIRequest() *AlibabaIdleRentOrderLogisticsDeliverAPIRequest {
	return poolAlibabaIdleRentOrderLogisticsDeliverAPIRequest.Get().(*AlibabaIdleRentOrderLogisticsDeliverAPIRequest)
}

// ReleaseAlibabaIdleRentOrderLogisticsDeliverAPIRequest 将 AlibabaIdleRentOrderLogisticsDeliverAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentOrderLogisticsDeliverAPIRequest(v *AlibabaIdleRentOrderLogisticsDeliverAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentOrderLogisticsDeliverAPIRequest.Put(v)
}
