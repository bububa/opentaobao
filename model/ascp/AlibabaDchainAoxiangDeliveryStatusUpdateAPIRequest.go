package ascp

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest 启用/停用配资源 API请求
// alibaba.dchain.aoxiang.delivery.status.update
//
// 启用/停用配资源
type AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest struct {
	model.Params
	// 启用/停用配资源入参
	_deliveryStatusUpdateRequest *DeliveryStatusUpdateRequest
}

// NewAlibabaDchainAoxiangDeliveryStatusUpdateRequest 初始化AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest对象
func NewAlibabaDchainAoxiangDeliveryStatusUpdateRequest() *AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest {
	return &AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) Reset() {
	r._deliveryStatusUpdateRequest = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.delivery.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryStatusUpdateRequest is DeliveryStatusUpdateRequest Setter
// 启用/停用配资源入参
func (r *AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) SetDeliveryStatusUpdateRequest(_deliveryStatusUpdateRequest *DeliveryStatusUpdateRequest) error {
	r._deliveryStatusUpdateRequest = _deliveryStatusUpdateRequest
	r.Set("delivery_status_update_request", _deliveryStatusUpdateRequest)
	return nil
}

// GetDeliveryStatusUpdateRequest DeliveryStatusUpdateRequest Getter
func (r AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) GetDeliveryStatusUpdateRequest() *DeliveryStatusUpdateRequest {
	return r._deliveryStatusUpdateRequest
}

var poolAlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaDchainAoxiangDeliveryStatusUpdateRequest()
	},
}

// GetAlibabaDchainAoxiangDeliveryStatusUpdateRequest 从 sync.Pool 获取 AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest
func GetAlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest() *AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest {
	return poolAlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest.Get().(*AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest)
}

// ReleaseAlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest 将 AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest 放入 sync.Pool
func ReleaseAlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest(v *AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) {
	v.Reset()
	poolAlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest.Put(v)
}
