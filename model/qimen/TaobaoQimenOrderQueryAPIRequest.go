package qimen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoqimenorderqueryAPIRequest 根据收件人信息查询交易单号接口 API请求
// taobao.qimen.order.query
//
// WMS 调用该接口，根据收件人信息查询平台交易订单号。
type TaobaoqimenorderqueryAPIRequest struct {
	model.Params
	// request
	_request *TaobaoqimenorderqueryRequest
}

// NewTaobaoqimenorderqueryRequest 初始化TaobaoqimenorderqueryAPIRequest对象
func NewTaobaoqimenorderqueryRequest() *TaobaoqimenorderqueryAPIRequest {
	return &TaobaoqimenorderqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoqimenorderqueryAPIRequest) GetApiMethodName() string {
	return "taobao.qimen.order.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoqimenorderqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoqimenorderqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequest is Request Setter
// request
func (r *TaobaoqimenorderqueryAPIRequest) SetRequest(_request *TaobaoqimenorderqueryRequest) error {
	r._request = _request
	r.Set("request", _request)
	return nil
}

// GetRequest Request Getter
func (r TaobaoqimenorderqueryAPIRequest) GetRequest() *TaobaoqimenorderqueryRequest {
	return r._request
}
