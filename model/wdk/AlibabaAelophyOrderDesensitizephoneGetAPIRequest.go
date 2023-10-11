package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderDesensitizephoneGetAPIRequest 获取订单隐私号 API请求
// alibaba.aelophy.order.desensitizephone.get
//
// 获取订单隐私号
type AlibabaAelophyOrderDesensitizephoneGetAPIRequest struct {
	model.Params
	// 请求
	_orderDesensitizePhoneRequest *OrderDesensitizePhoneRequest
}

// NewAlibabaAelophyOrderDesensitizephoneGetRequest 初始化AlibabaAelophyOrderDesensitizephoneGetAPIRequest对象
func NewAlibabaAelophyOrderDesensitizephoneGetRequest() *AlibabaAelophyOrderDesensitizephoneGetAPIRequest {
	return &AlibabaAelophyOrderDesensitizephoneGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyOrderDesensitizephoneGetAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.desensitizephone.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyOrderDesensitizephoneGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyOrderDesensitizephoneGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderDesensitizePhoneRequest is OrderDesensitizePhoneRequest Setter
// 请求
func (r *AlibabaAelophyOrderDesensitizephoneGetAPIRequest) SetOrderDesensitizePhoneRequest(_orderDesensitizePhoneRequest *OrderDesensitizePhoneRequest) error {
	r._orderDesensitizePhoneRequest = _orderDesensitizePhoneRequest
	r.Set("order_desensitize_phone_request", _orderDesensitizePhoneRequest)
	return nil
}

// GetOrderDesensitizePhoneRequest OrderDesensitizePhoneRequest Getter
func (r AlibabaAelophyOrderDesensitizephoneGetAPIRequest) GetOrderDesensitizePhoneRequest() *OrderDesensitizePhoneRequest {
	return r._orderDesensitizePhoneRequest
}
