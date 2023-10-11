package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAelophyOrderDelivererChangeAPIRequest 配送员信息变更接口 API请求
// alibaba.aelophy.order.deliverer.change
//
// 配送员信息变更接口
type AlibabaAelophyOrderDelivererChangeAPIRequest struct {
	model.Params
	// 配送员信息变更请求
	_delivererChangeRequest *DelivererChangeRequest
}

// NewAlibabaAelophyOrderDelivererChangeRequest 初始化AlibabaAelophyOrderDelivererChangeAPIRequest对象
func NewAlibabaAelophyOrderDelivererChangeRequest() *AlibabaAelophyOrderDelivererChangeAPIRequest {
	return &AlibabaAelophyOrderDelivererChangeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetApiMethodName() string {
	return "alibaba.aelophy.order.deliverer.change"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetDelivererChangeRequest is DelivererChangeRequest Setter
// 配送员信息变更请求
func (r *AlibabaAelophyOrderDelivererChangeAPIRequest) SetDelivererChangeRequest(_delivererChangeRequest *DelivererChangeRequest) error {
	r._delivererChangeRequest = _delivererChangeRequest
	r.Set("deliverer_change_request", _delivererChangeRequest)
	return nil
}

// GetDelivererChangeRequest DelivererChangeRequest Getter
func (r AlibabaAelophyOrderDelivererChangeAPIRequest) GetDelivererChangeRequest() *DelivererChangeRequest {
	return r._delivererChangeRequest
}
