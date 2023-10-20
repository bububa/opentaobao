package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderSenditemAPIRequest 确认发货 API请求
// alibaba.idle.rent.order.senditem
//
// 确认发货
type AlibabaIdleRentOrderSenditemAPIRequest struct {
	model.Params
	// 物流信息
	_logisticsList []LogisticsDto
	// 订单id
	_orderId int64
}

// NewAlibabaIdleRentOrderSenditemRequest 初始化AlibabaIdleRentOrderSenditemAPIRequest对象
func NewAlibabaIdleRentOrderSenditemRequest() *AlibabaIdleRentOrderSenditemAPIRequest {
	return &AlibabaIdleRentOrderSenditemAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentOrderSenditemAPIRequest) Reset() {
	r._logisticsList = r._logisticsList[:0]
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.senditem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLogisticsList is LogisticsList Setter
// 物流信息
func (r *AlibabaIdleRentOrderSenditemAPIRequest) SetLogisticsList(_logisticsList []LogisticsDto) error {
	r._logisticsList = _logisticsList
	r.Set("logistics_list", _logisticsList)
	return nil
}

// GetLogisticsList LogisticsList Getter
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetLogisticsList() []LogisticsDto {
	return r._logisticsList
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderSenditemAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaIdleRentOrderSenditemAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlibabaIdleRentOrderSenditemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentOrderSenditemRequest()
	},
}

// GetAlibabaIdleRentOrderSenditemRequest 从 sync.Pool 获取 AlibabaIdleRentOrderSenditemAPIRequest
func GetAlibabaIdleRentOrderSenditemAPIRequest() *AlibabaIdleRentOrderSenditemAPIRequest {
	return poolAlibabaIdleRentOrderSenditemAPIRequest.Get().(*AlibabaIdleRentOrderSenditemAPIRequest)
}

// ReleaseAlibabaIdleRentOrderSenditemAPIRequest 将 AlibabaIdleRentOrderSenditemAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentOrderSenditemAPIRequest(v *AlibabaIdleRentOrderSenditemAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentOrderSenditemAPIRequest.Put(v)
}
