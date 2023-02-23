package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenDeliveryorderCreateAPIRequest 发货单创建接口 API请求
// taobao.qimen.deliveryorder.create
//
// taobao.qimen.deliveryorder.create
type TaobaoQimenDeliveryorderCreateAPIRequest struct {
	model.Params
	// 奇门仓储字段
	_request *DeliveryOrderCreateRequest
}

// NewTaobaoQimenDeliveryorderCreateRequest 初始化TaobaoQimenDeliveryorderCreateAPIRequest对象
func NewTaobaoQimenDeliveryorderCreateRequest() *TaobaoQimenDeliveryorderCreateAPIRequest {
	return &TaobaoQimenDeliveryorderCreateAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.deliveryorder.create"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// 奇门仓储字段
func (r *TaobaoQimenDeliveryorderCreateAPIRequest) SetRequest(_request *DeliveryOrderCreateRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenDeliveryorderCreateAPIRequest) GetRequest() *DeliveryOrderCreateRequest {
	return r._request
}
