package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoQimenDeliveryorderConfirmAPIRequest
发货单确认接口 API请求
taobao.qimen.deliveryorder.confirm

taobao.qimen.deliveryorder.confirm */
type TaobaoQimenDeliveryorderConfirmAPIRequest struct {
	model.Params
	//
	_request *DeliveryOrderConfirmRequest
}

// NewTaobaoQimenDeliveryorderConfirmRequest 初始化TaobaoQimenDeliveryorderConfirmAPIRequest对象
func NewTaobaoQimenDeliveryorderConfirmRequest() *TaobaoQimenDeliveryorderConfirmAPIRequest {
	return &TaobaoQimenDeliveryorderConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenDeliveryorderConfirmAPIRequest) SetRequest(_request *DeliveryOrderConfirmRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenDeliveryorderConfirmAPIRequest) GetRequest() *DeliveryOrderConfirmRequest {
	return r._request
}
