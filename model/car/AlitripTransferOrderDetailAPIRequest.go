package car

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripTransferOrderDetailAPIRequest 接送订单详情接口 API请求
// alitrip.transfer.order.detail
//
// 接送订单详情接口
type AlitripTransferOrderDetailAPIRequest struct {
	model.Params
	// 飞猪平台订单id
	_orderId string
	// 分配的服务商Id
	_providerId string
}

// NewAlitripTransferOrderDetailRequest 初始化AlitripTransferOrderDetailAPIRequest对象
func NewAlitripTransferOrderDetailRequest() *AlitripTransferOrderDetailAPIRequest {
	return &AlitripTransferOrderDetailAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlitripTransferOrderDetailAPIRequest) Reset() {
	r._orderId = ""
	r._providerId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlitripTransferOrderDetailAPIRequest) GetApiMethodName() string {
	return "alitrip.transfer.order.detail"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlitripTransferOrderDetailAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlitripTransferOrderDetailAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 飞猪平台订单id
func (r *AlitripTransferOrderDetailAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlitripTransferOrderDetailAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetProviderId is ProviderId Setter
// 分配的服务商Id
func (r *AlitripTransferOrderDetailAPIRequest) SetProviderId(_providerId string) error {
	r._providerId = _providerId
	r.Set("provider_id", _providerId)
	return nil
}

// GetProviderId ProviderId Getter
func (r AlitripTransferOrderDetailAPIRequest) GetProviderId() string {
	return r._providerId
}

var poolAlitripTransferOrderDetailAPIRequest = sync.Pool{
	New: func() any {
		return NewAlitripTransferOrderDetailRequest()
	},
}

// GetAlitripTransferOrderDetailRequest 从 sync.Pool 获取 AlitripTransferOrderDetailAPIRequest
func GetAlitripTransferOrderDetailAPIRequest() *AlitripTransferOrderDetailAPIRequest {
	return poolAlitripTransferOrderDetailAPIRequest.Get().(*AlitripTransferOrderDetailAPIRequest)
}

// ReleaseAlitripTransferOrderDetailAPIRequest 将 AlitripTransferOrderDetailAPIRequest 放入 sync.Pool
func ReleaseAlitripTransferOrderDetailAPIRequest(v *AlitripTransferOrderDetailAPIRequest) {
	v.Reset()
	poolAlitripTransferOrderDetailAPIRequest.Put(v)
}
