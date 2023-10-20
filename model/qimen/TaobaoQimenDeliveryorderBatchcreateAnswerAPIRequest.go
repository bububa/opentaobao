package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryorderbatchcreateanswerAPIRequest 发货单创建结果通知接口(批量) API请求
// taobao.qimen.deliveryorder.batchcreate.answer
//
// WMS调用接口，用于异步化的批量发货单创建结果通知。（如菜鸟发货单批量创建结果的返回）
type TaobaoqimendeliveryorderbatchcreateanswerAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchCreateAnswerRequest
}

// NewTaobaoqimendeliveryorderbatchcreateanswerRequest 初始化TaobaoqimendeliveryorderbatchcreateanswerAPIRequest对象
func NewTaobaoqimendeliveryorderbatchcreateanswerRequest() *TaobaoqimendeliveryorderbatchcreateanswerAPIRequest {
	return &TaobaoqimendeliveryorderbatchcreateanswerAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimendeliveryorderbatchcreateanswerAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchcreate.answer"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimendeliveryorderbatchcreateanswerAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimendeliveryorderbatchcreateanswerAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimendeliveryorderbatchcreateanswerAPIRequest) SetRequest(_request *DeliveryOrderBatchCreateAnswerRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimendeliveryorderbatchcreateanswerAPIRequest) GetRequest() *DeliveryOrderBatchCreateAnswerRequest {
	return r._request
}
