package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderprocessQueryAPIRequest 订单流水查询接口 API请求
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
type TaobaoQimenOrderprocessQueryAPIRequest struct {
	model.Params
	//
	_request *OrderProcessQueryRequest
}

// NewTaobaoQimenOrderprocessQueryRequest 初始化TaobaoQimenOrderprocessQueryAPIRequest对象
func NewTaobaoQimenOrderprocessQueryRequest() *TaobaoQimenOrderprocessQueryAPIRequest {
	return &TaobaoQimenOrderprocessQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoQimenOrderprocessQueryAPIRequest) SetRequest(_request *OrderProcessQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoQimenOrderprocessQueryAPIRequest) GetRequest() *OrderProcessQueryRequest {
	return r._request
}
