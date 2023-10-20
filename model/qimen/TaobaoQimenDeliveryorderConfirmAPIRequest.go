package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimendeliveryorderconfirmAPIRequest 发货单确认接口 API请求
// taobao.qimen.deliveryorder.confirm
//
// taobao.qimen.deliveryorder.confirm
type TaobaoqimendeliveryorderconfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderConfirmRequest
}

// NewTaobaoqimendeliveryorderconfirmRequest 初始化TaobaoqimendeliveryorderconfirmAPIRequest对象
func NewTaobaoqimendeliveryorderconfirmRequest() *TaobaoqimendeliveryorderconfirmAPIRequest {
	return &TaobaoqimendeliveryorderconfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimendeliveryorderconfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimendeliveryorderconfirmAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimendeliveryorderconfirmAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimendeliveryorderconfirmAPIRequest) SetRequest(_request *DeliveryOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimendeliveryorderconfirmAPIRequest) GetRequest() *DeliveryOrderConfirmRequest {
	return r._request
}
