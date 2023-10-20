package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest 新建/更新配资源 API请求
// alibaba.dchain.aoxiang.delivery.create.update
//
// 新建/更新配资源
type AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest struct {
	model.Params
	// 新建/更新配资源入参
	_deliveryUpsertRequest *DeliveryUpsertRequest
}

// NewAlibabaDchainAoxiangDeliveryCreateUpdateRequest 初始化AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest对象
func NewAlibabaDchainAoxiangDeliveryCreateUpdateRequest() *AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest {
	return &AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest) GetApiMethodName() string {
	return "alibaba.dchain.aoxiang.delivery.create.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryUpsertRequest is DeliveryUpsertRequest Setter
// 新建/更新配资源入参
func (r *AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest) SetDeliveryUpsertRequest(_deliveryUpsertRequest *DeliveryUpsertRequest) error {
	r._deliveryUpsertRequest = _deliveryUpsertRequest
	r.Set("delivery_upsert_request", _deliveryUpsertRequest)
	return nil
}

// GetDeliveryUpsertRequest DeliveryUpsertRequest Getter
func (r AlibabaDchainAoxiangDeliveryCreateUpdateAPIRequest) GetDeliveryUpsertRequest() *DeliveryUpsertRequest {
	return r._deliveryUpsertRequest
}
