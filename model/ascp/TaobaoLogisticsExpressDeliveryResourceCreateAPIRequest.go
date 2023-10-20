package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsexpressdeliveryresourcecreateAPIRequest 新建/更新配资源 API请求
// taobao.logistics.express.delivery.resource.create
//
// 新建/更新配资源
type TaobaologisticsexpressdeliveryresourcecreateAPIRequest struct {
	model.Params
	// 配资源
	_deliveryResourceCreateRequest *DeliveryResourceCreateRequest
}

// NewTaobaologisticsexpressdeliveryresourcecreateRequest 初始化TaobaologisticsexpressdeliveryresourcecreateAPIRequest对象
func NewTaobaologisticsexpressdeliveryresourcecreateRequest() *TaobaologisticsexpressdeliveryresourcecreateAPIRequest {
	return &TaobaologisticsexpressdeliveryresourcecreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsexpressdeliveryresourcecreateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.express.delivery.resource.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsexpressdeliveryresourcecreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsexpressdeliveryresourcecreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryResourceCreateRequest is DeliveryResourceCreateRequest Setter
// 配资源
func (r *TaobaologisticsexpressdeliveryresourcecreateAPIRequest) SetDeliveryResourceCreateRequest(_deliveryResourceCreateRequest *DeliveryResourceCreateRequest) error {
	r._deliveryResourceCreateRequest = _deliveryResourceCreateRequest
	r.Set("delivery_resource_create_request", _deliveryResourceCreateRequest)
	return nil
}

// GetDeliveryResourceCreateRequest DeliveryResourceCreateRequest Getter
func (r TaobaologisticsexpressdeliveryresourcecreateAPIRequest) GetDeliveryResourceCreateRequest() *DeliveryResourceCreateRequest {
	return r._deliveryResourceCreateRequest
}
