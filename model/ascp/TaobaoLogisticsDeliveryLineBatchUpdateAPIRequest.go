package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsdeliverylinebatchupdateAPIRequest 线路能力创建/更新 API请求
// taobao.logistics.delivery.line.batch.update
//
// 线路能力创建/更新
type TaobaologisticsdeliverylinebatchupdateAPIRequest struct {
	model.Params
	// 线路能力创建/更新入参
	_deliveryLineBatchUpdateRequest *DeliveryLineBatchUpdateRequest
}

// NewTaobaologisticsdeliverylinebatchupdateRequest 初始化TaobaologisticsdeliverylinebatchupdateAPIRequest对象
func NewTaobaologisticsdeliverylinebatchupdateRequest() *TaobaologisticsdeliverylinebatchupdateAPIRequest {
	return &TaobaologisticsdeliverylinebatchupdateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsdeliverylinebatchupdateAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.delivery.line.batch.update"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsdeliverylinebatchupdateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsdeliverylinebatchupdateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryLineBatchUpdateRequest is DeliveryLineBatchUpdateRequest Setter
// 线路能力创建/更新入参
func (r *TaobaologisticsdeliverylinebatchupdateAPIRequest) SetDeliveryLineBatchUpdateRequest(_deliveryLineBatchUpdateRequest *DeliveryLineBatchUpdateRequest) error {
	r._deliveryLineBatchUpdateRequest = _deliveryLineBatchUpdateRequest
	r.Set("delivery_line_batch_update_request", _deliveryLineBatchUpdateRequest)
	return nil
}

// GetDeliveryLineBatchUpdateRequest DeliveryLineBatchUpdateRequest Getter
func (r TaobaologisticsdeliverylinebatchupdateAPIRequest) GetDeliveryLineBatchUpdateRequest() *DeliveryLineBatchUpdateRequest {
	return r._deliveryLineBatchUpdateRequest
}
