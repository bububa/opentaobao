package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryorderbatchconfirmAPIRequest 发货单确认接口 API请求
// taobao.qimen.deliveryorder.batchconfirm
//
// taobao.qimen.deliveryorder.batchconfirm
type TaobaoqimendeliveryorderbatchconfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderBatchConfirmRequest
}

// NewTaobaoqimendeliveryorderbatchconfirmRequest 初始化TaobaoqimendeliveryorderbatchconfirmAPIRequest对象
func NewTaobaoqimendeliveryorderbatchconfirmRequest() *TaobaoqimendeliveryorderbatchconfirmAPIRequest {
	return &TaobaoqimendeliveryorderbatchconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimendeliveryorderbatchconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.batchconfirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimendeliveryorderbatchconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimendeliveryorderbatchconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimendeliveryorderbatchconfirmAPIRequest) SetRequest(_request *DeliveryOrderBatchConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimendeliveryorderbatchconfirmAPIRequest) GetRequest() *DeliveryOrderBatchConfirmRequest {
	return r._request
}
