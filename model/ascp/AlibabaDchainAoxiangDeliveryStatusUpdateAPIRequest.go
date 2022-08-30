package ascp

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.delivery.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangDeliveryStatusUpdateAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
