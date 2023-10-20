package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchcreateAPIRequest 发货单创建批量接口 API请求
// taobao.qimen.deliveryorder.batchcreate
//
// taobao.qimen.deliveryorder.batchcreate
type TaobaoQimenDeliveryorderBatchcreateAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateRequest
}

// NewTaobaoQimenDeliveryorderBatchcreateRequest 初始化TaobaoQimenDeliveryorderBatchcreateAPIRequest对象
func NewTaobaoQimenDeliveryorderBatchcreateRequest() *TaobaoQimenDeliveryorderBatchcreateAPIRequest {
	return &TaobaoQimenDeliveryorderBatchcreateAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenDeliveryorderBatchcreateAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchcreate"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenDeliveryorderBatchcreateAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderBatchcreateAPIRequest) GetRequest() *DeliveryOrderBatchCreateRequest {
	return r._request
}

var poolTaobaoQimenDeliveryorderBatchcreateAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenDeliveryorderBatchcreateRequest()
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateRequest 从 sync.Pool 获取 TaobaoQimenDeliveryorderBatchcreateAPIRequest
func GetTaobaoQimenDeliveryorderBatchcreateAPIRequest() *TaobaoQimenDeliveryorderBatchcreateAPIRequest {
	return poolTaobaoQimenDeliveryorderBatchcreateAPIRequest.Get().(*TaobaoQimenDeliveryorderBatchcreateAPIRequest)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateAPIRequest 将 TaobaoQimenDeliveryorderBatchcreateAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenDeliveryorderBatchcreateAPIRequest(v *TaobaoQimenDeliveryorderBatchcreateAPIRequest) {
	v.Reset()
	poolTaobaoQimenDeliveryorderBatchcreateAPIRequest.Put(v)
}
