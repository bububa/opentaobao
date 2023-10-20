package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaelophyorderdesensitizephonegetAPIRequest 获取订单隐私号 API请求
// alibaba.aelophy.order.desensitizephone.get
//
// 获取订单隐私号
type AlibabaaelophyorderdesensitizephonegetAPIRequest struct {
	model.Params
	// 请求
	_orderDesensitizePhoneRequest *OrderDesensitizePhoneRequest
}

// NewAlibabaaelophyorderdesensitizephonegetRequest 初始化AlibabaaelophyorderdesensitizephonegetAPIRequest对象
func NewAlibabaaelophyorderdesensitizephonegetRequest() *AlibabaaelophyorderdesensitizephonegetAPIRequest {
	return &AlibabaaelophyorderdesensitizephonegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaaelophyorderdesensitizephonegetAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.desensitizephone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaaelophyorderdesensitizephonegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaaelophyorderdesensitizephonegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderDesensitizePhoneRequest is OrderDesensitizePhoneRequest Setter
// 请求
func (r *AlibabaaelophyorderdesensitizephonegetAPIRequest) SetOrderDesensitizePhoneRequest(_orderDesensitizePhoneRequest *OrderDesensitizePhoneRequest) error {
	r._orderDesensitizePhoneRequest = _orderDesensitizePhoneRequest
	r.Set("order_desensitize_phone_request", _orderDesensitizePhoneRequest)
	return nil
}

// GetOrderDesensitizePhoneRequest OrderDesensitizePhoneRequest Getter
func (r AlibabaaelophyorderdesensitizephonegetAPIRequest) GetOrderDesensitizePhoneRequest() *OrderDesensitizePhoneRequest {
	return r._orderDesensitizePhoneRequest
}
