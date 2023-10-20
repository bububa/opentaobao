package idle

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaIdleRentOrderPackageAPIRequest) Reset() {
	r._orderId = 0
	r._logistics = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIdleRentOrderPackageAPIRequest) GetApiMethodName() string {
	return "alibaba.idle.rent.order.package"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIdleRentOrderPackageAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIdleRentOrderPackageAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单id
func (r *AlibabaIdleRentOrderPackageAPIRequest) SetOrderId(_orderId int64) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaIdleRentOrderPackageAPIRequest) GetOrderId() int64 {
	return r._orderId
}

// SetLogistics is Logistics Setter
// 物流信息
func (r *AlibabaIdleRentOrderPackageAPIRequest) SetLogistics(_logistics *LogisticsDto) error {
	r._logistics = _logistics
	r.Set("logistics", _logistics)
	return nil
}

// GetLogistics Logistics Getter
func (r AlibabaIdleRentOrderPackageAPIRequest) GetLogistics() *LogisticsDto {
	return r._logistics
}

var poolAlibabaIdleRentOrderPackageAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaIdleRentOrderPackageRequest()
	},
}

// GetAlibabaIdleRentOrderPackageRequest 从 sync.Pool 获取 AlibabaIdleRentOrderPackageAPIRequest
func GetAlibabaIdleRentOrderPackageAPIRequest() *AlibabaIdleRentOrderPackageAPIRequest {
	return poolAlibabaIdleRentOrderPackageAPIRequest.Get().(*AlibabaIdleRentOrderPackageAPIRequest)
}

// ReleaseAlibabaIdleRentOrderPackageAPIRequest 将 AlibabaIdleRentOrderPackageAPIRequest 放入 sync.Pool
func ReleaseAlibabaIdleRentOrderPackageAPIRequest(v *AlibabaIdleRentOrderPackageAPIRequest) {
	v.Reset()
	poolAlibabaIdleRentOrderPackageAPIRequest.Put(v)
}
