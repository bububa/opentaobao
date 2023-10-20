package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabadchainaoxiangdeliverystatusupdateAPIRequest 启用/停用配资源 API请求
// alibaba.dchain.aoxiang.delivery.status.update
//
// 启用/停用配资源
type AlibabadchainaoxiangdeliverystatusupdateAPIRequest struct {
	model.Params
	// 启用/停用配资源入参
	_deliveryStatusUpdateRequest *DeliveryStatusUpdateRequest
}

// NewAlibabadchainaoxiangdeliverystatusupdateRequest 初始化AlibabadchainaoxiangdeliverystatusupdateAPIRequest对象
func NewAlibabadchainaoxiangdeliverystatusupdateRequest() *AlibabadchainaoxiangdeliverystatusupdateAPIRequest {
	return &AlibabadchainaoxiangdeliverystatusupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabadchainaoxiangdeliverystatusupdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.delivery.status.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabadchainaoxiangdeliverystatusupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabadchainaoxiangdeliverystatusupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryStatusUpdateRequest is DeliveryStatusUpdateRequest Setter
// 启用/停用配资源入参
func (r *AlibabadchainaoxiangdeliverystatusupdateAPIRequest) SetDeliveryStatusUpdateRequest(_deliveryStatusUpdateRequest *DeliveryStatusUpdateRequest) error {
	r._deliveryStatusUpdateRequest = _deliveryStatusUpdateRequest
	r.Set("delivery_status_update_request", _deliveryStatusUpdateRequest)
	return nil
}

// GetDeliveryStatusUpdateRequest DeliveryStatusUpdateRequest Getter
func (r AlibabadchainaoxiangdeliverystatusupdateAPIRequest) GetDeliveryStatusUpdateRequest() *DeliveryStatusUpdateRequest {
	return r._deliveryStatusUpdateRequest
}
