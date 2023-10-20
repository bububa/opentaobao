package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest 线路能力删除 API请求
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
type TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest struct {
	model.Params
	// 线路能力删除入参
	_deliveryLineBatchDeleteRequest *DeliveryLineBatchDeleteRequest
}

// NewTaobaoLogisticsDeliveryLineBatchDeleteRequest 初始化TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest对象
func NewTaobaoLogisticsDeliveryLineBatchDeleteRequest() *TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest {
	return &TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.delivery.line.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryLineBatchDeleteRequest is DeliveryLineBatchDeleteRequest Setter
// 线路能力删除入参
func (r *TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) SetDeliveryLineBatchDeleteRequest(_deliveryLineBatchDeleteRequest *DeliveryLineBatchDeleteRequest) error {
	r._deliveryLineBatchDeleteRequest = _deliveryLineBatchDeleteRequest
	r.Set("delivery_line_batch_delete_request", _deliveryLineBatchDeleteRequest)
	return nil
}

// GetDeliveryLineBatchDeleteRequest DeliveryLineBatchDeleteRequest Getter
func (r TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) GetDeliveryLineBatchDeleteRequest() *DeliveryLineBatchDeleteRequest {
	return r._deliveryLineBatchDeleteRequest
}
