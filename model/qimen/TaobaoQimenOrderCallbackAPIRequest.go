package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoQimenOrderCallbackAPIRequest 配送拦截接口 API请求
// taobao.qimen.order.callback
//
// 配送拦截
type TaobaoQimenOrderCallbackAPIRequest struct {
	model.Params
	//
	_request *OrderCallbackRequestDo
}

// NewTaobaoQimenOrderCallbackRequest 初始化TaobaoQimenOrderCallbackAPIRequest对象
func NewTaobaoQimenOrderCallbackRequest() *TaobaoQimenOrderCallbackAPIRequest {
	return &TaobaoQimenOrderCallbackAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoQimenOrderCallbackAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.callback"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoQimenOrderCallbackAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Request Setter
//
func (r *TaobaoQimenOrderCallbackAPIRequest) SetRequest(_request *OrderCallbackRequestDo) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// Get Request Getter
func (r TaobaoQimenOrderCallbackAPIRequest) GetRequest() *OrderCallbackRequestDo {
	return r._request
}
