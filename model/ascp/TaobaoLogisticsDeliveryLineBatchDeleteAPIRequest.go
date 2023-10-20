package ascp

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) Reset() {
	r._deliveryLineBatchDeleteRequest = nil
	r.Params.ToZero()
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

var poolTaobaoLogisticsDeliveryLineBatchDeleteAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoLogisticsDeliveryLineBatchDeleteRequest()
	},
}

// GetTaobaoLogisticsDeliveryLineBatchDeleteRequest 从 sync.Pool 获取 TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest
func GetTaobaoLogisticsDeliveryLineBatchDeleteAPIRequest() *TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest {
	return poolTaobaoLogisticsDeliveryLineBatchDeleteAPIRequest.Get().(*TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest)
}

// ReleaseTaobaoLogisticsDeliveryLineBatchDeleteAPIRequest 将 TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest 放入 sync.Pool
func ReleaseTaobaoLogisticsDeliveryLineBatchDeleteAPIRequest(v *TaobaoLogisticsDeliveryLineBatchDeleteAPIRequest) {
	v.Reset()
	poolTaobaoLogisticsDeliveryLineBatchDeleteAPIRequest.Put(v)
}
