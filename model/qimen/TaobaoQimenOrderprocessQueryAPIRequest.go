package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderprocessqueryAPIRequest 订单流水查询接口 API请求
// taobao.qimen.orderprocess.query
//
// ERP调用订单流水查询接口
type TaobaoqimenorderprocessqueryAPIRequest struct {
	model.Params
	//
	_request *OrderProcessQueryRequest
}

// NewTaobaoqimenorderprocessqueryRequest 初始化TaobaoqimenorderprocessqueryAPIRequest对象
func NewTaobaoqimenorderprocessqueryRequest() *TaobaoqimenorderprocessqueryAPIRequest {
	return &TaobaoqimenorderprocessqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenorderprocessqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.orderprocess.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenorderprocessqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenorderprocessqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenorderprocessqueryAPIRequest) SetRequest(_request *OrderProcessQueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenorderprocessqueryAPIRequest) GetRequest() *OrderProcessQueryRequest {
	return r._request
}
