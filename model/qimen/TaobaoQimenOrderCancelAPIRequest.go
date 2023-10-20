package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenordercancelAPIRequest 单据取消接口 API请求
// taobao.qimen.order.cancel
//
// taobao.qimen.order.cancel
type TaobaoqimenordercancelAPIRequest struct {
	model.Params
	//
	_request *OrderCancelRequest
}

// NewTaobaoqimenordercancelRequest 初始化TaobaoqimenordercancelAPIRequest对象
func NewTaobaoqimenordercancelRequest() *TaobaoqimenordercancelAPIRequest {
	return &TaobaoqimenordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenordercancelAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
func (r *TaobaoqimenordercancelAPIRequest) SetRequest(_request *OrderCancelRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenordercancelAPIRequest) GetRequest() *OrderCancelRequest {
	return r._request
}
