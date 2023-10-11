package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest 新建/更新配资源 API请求
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
type TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest struct {
	model.Params
	// 配资源
	_deliveryResourceCreateRequest *DeliveryResourceCreateRequest
}

// NewTaobaoLogisticsExpressDeliveryResourceCreateRequest 初始化TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest对象
func NewTaobaoLogisticsExpressDeliveryResourceCreateRequest() *TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest {
	return &TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.resource.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryResourceCreateRequest is DeliveryResourceCreateRequest Setter
// 配资源
func (r *TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest) SetDeliveryResourceCreateRequest(_deliveryResourceCreateRequest *DeliveryResourceCreateRequest) error {
	r._deliveryResourceCreateRequest = _deliveryResourceCreateRequest
	r.Set("delivery_resource_create_request", _deliveryResourceCreateRequest)
	return nil
}

// GetDeliveryResourceCreateRequest DeliveryResourceCreateRequest Getter
func (r TaobaoLogisticsExpressDeliveryResourceCreateAPIRequest) GetDeliveryResourceCreateRequest() *DeliveryResourceCreateRequest {
	return r._deliveryResourceCreateRequest
}
