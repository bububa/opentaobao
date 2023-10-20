package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenordercallbackAPIRequest 配送拦截接口 API请求
// taobao.qimen.order.callback
//
// 配送拦截
type TaobaoqimenordercallbackAPIRequest struct {
	model.Params
	//
	_request *OrderCallbackRequestDo
}

// NewTaobaoqimenordercallbackRequest 初始化TaobaoqimenordercallbackAPIRequest对象
func NewTaobaoqimenordercallbackRequest() *TaobaoqimenordercallbackAPIRequest {
	return &TaobaoqimenordercallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenordercallbackAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenordercallbackAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenordercallbackAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenordercallbackAPIRequest) SetRequest(_request *OrderCallbackRequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenordercallbackAPIRequest) GetRequest() *OrderCallbackRequestDo {
	return r._request
}
