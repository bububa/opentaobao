package qimen

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest 发货单创建结果通知接口(批量) API请求
// taobao.qimen.deliveryorder.batchcreate.answer
//
// WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateAnswerRequest
}

// NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest 初始化TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest对象
func NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest() *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest {
	return &TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) Reset() {
	r._request = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchcreate.answer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateAnswerRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetRequest() *DeliveryOrderBatchCreateAnswerRequest {
	return r._request
}

var poolTaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest()
	},
}

// GetTaobaoQimenDeliveryorderBatchcreateAnswerRequest 从 sync.Pool 获取 TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest
func GetTaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest() *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest {
	return poolTaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest.Get().(*TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest)
}

// ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest 将 TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest 放入 sync.Pool
func ReleaseTaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest(v *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) {
	v.Reset()
	poolTaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest.Put(v)
}
