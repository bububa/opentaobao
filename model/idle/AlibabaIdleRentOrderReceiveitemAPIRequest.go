package idle

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleRentOrderReceiveitemAPIRequest 确认签收 API请求
// alibaba.idle.rent.order.receiveitem
//
// 确认揽收/签收
type AlibabaIdleRentOrderReceiveitemAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// NewAlibabaIdleRentOrderReceiveitemRequest 初始化AlibabaIdleRentOrderReceiveitemAPIRequest对象
func NewAlibabaIdleRentOrderReceiveitemRequest() *AlibabaIdleRentOrderReceiveitemAPIRequest {
	return &AlibabaIdleRentOrderReceiveitemAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentOrderReceiveitemAPIRequest) Reset() {
	r._orderId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.receiveitem"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderReceiveitemAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaIdleRentOrderReceiveitemAPIRequest) GetOrderId() int64 {
	return r._orderId
}

var poolAlibabaIdleRentOrderReceiveitemAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentOrderReceiveitemRequest()
	},
}

// GetAlibabaIdleRentOrderReceiveitemRequest 从 sync.Pool 获取 AlibabaIdleRentOrderReceiveitemAPIRequest
func GetAlibabaIdleRentOrderReceiveitemAPIRequest() *AlibabaIdleRentOrderReceiveitemAPIRequest {
	return poolAlibabaIdleRentOrderReceiveitemAPIRequest.Get().(*AlibabaIdleRentOrderReceiveitemAPIRequest)
}

// ReleaseAlibabaIdleRentOrderReceiveitemAPIRequest 将 AlibabaIdleRentOrderReceiveitemAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentOrderReceiveitemAPIRequest(v *AlibabaIdleRentOrderReceiveitemAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentOrderReceiveitemAPIRequest.Put(v)
}
