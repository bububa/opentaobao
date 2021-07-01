package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest
发货单创建结果通知接口(批量) API请求
taobao.qimen.deliveryorder.batchcreate.answer

WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回） */
type TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateAnswerRequest
}

// NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest 初始化TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest对象
func NewTaobaoQimenDeliveryorderBatchcreateAnswerRequest() *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest {
	return &TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchcreate.answer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateAnswerRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenDeliveryorderBatchcreateAnswerAPIRequest) GetRequest() *DeliveryOrderBatchCreateAnswerRequest {
	return r._request
}
