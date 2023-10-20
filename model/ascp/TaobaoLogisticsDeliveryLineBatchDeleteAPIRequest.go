package ascp

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaologisticsdeliverylinebatchdeleteAPIRequest 线路能力删除 API请求
// taobao.logistics.delivery.line.batch.delete
//
// 线路能力删除
type TaobaologisticsdeliverylinebatchdeleteAPIRequest struct {
	model.Params
	// 线路能力删除入参
	_deliveryLineBatchDeleteRequest *DeliveryLineBatchDeleteRequest
}

// NewTaobaologisticsdeliverylinebatchdeleteRequest 初始化TaobaologisticsdeliverylinebatchdeleteAPIRequest对象
func NewTaobaologisticsdeliverylinebatchdeleteRequest() *TaobaologisticsdeliverylinebatchdeleteAPIRequest {
	return &TaobaologisticsdeliverylinebatchdeleteAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaologisticsdeliverylinebatchdeleteAPIRequest) GetApiMethodName() string {
	return "taobao.logistics.delivery.line.batch.delete"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaologisticsdeliverylinebatchdeleteAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaologisticsdeliverylinebatchdeleteAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDeliveryLineBatchDeleteRequest is DeliveryLineBatchDeleteRequest Setter
// 线路能力删除入参
func (r *TaobaologisticsdeliverylinebatchdeleteAPIRequest) SetDeliveryLineBatchDeleteRequest(_deliveryLineBatchDeleteRequest *DeliveryLineBatchDeleteRequest) error {
	r._deliveryLineBatchDeleteRequest = _deliveryLineBatchDeleteRequest
	r.Set("delivery_line_batch_delete_request", _deliveryLineBatchDeleteRequest)
	return nil
}

// GetDeliveryLineBatchDeleteRequest DeliveryLineBatchDeleteRequest Getter
func (r TaobaologisticsdeliverylinebatchdeleteAPIRequest) GetDeliveryLineBatchDeleteRequest() *DeliveryLineBatchDeleteRequest {
	return r._deliveryLineBatchDeleteRequest
}
