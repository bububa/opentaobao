package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoalitripbuyergetAPIRequest 敏感信息查询 API请求
// taobao.alitrip.buyer.get
//
// 针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
type TaobaoalitripbuyergetAPIRequest struct {
	model.Params
	// 敏感信息查询请求参数
	_requestAxb *RequestAxbDo
}

// NewTaobaoalitripbuyergetRequest 初始化TaobaoalitripbuyergetAPIRequest对象
func NewTaobaoalitripbuyergetRequest() *TaobaoalitripbuyergetAPIRequest {
	return &TaobaoalitripbuyergetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoalitripbuyergetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.buyer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoalitripbuyergetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoalitripbuyergetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestAxb is RequestAxb Setter
// 敏感信息查询请求参数
func (r *TaobaoalitripbuyergetAPIRequest) SetRequestAxb(_requestAxb *RequestAxbDo) error {
	r._requestAxb = _requestAxb
	r.Set("request_axb", _requestAxb)
	return nil
}

// GetRequestAxb RequestAxb Getter
func (r TaobaoalitripbuyergetAPIRequest) GetRequestAxb() *RequestAxbDo {
	return r._requestAxb
}
