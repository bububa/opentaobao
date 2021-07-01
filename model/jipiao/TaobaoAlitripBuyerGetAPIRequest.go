package jipiao

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripBuyerGetAPIRequest
敏感信息查询 API请求
taobao.alitrip.buyer.get

针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。 */
type TaobaoAlitripBuyerGetAPIRequest struct {
	model.Params
	// 敏感信息查询请求参数
	_requestAxb *RequestAxbDo
}

// NewTaobaoAlitripBuyerGetRequest 初始化TaobaoAlitripBuyerGetAPIRequest对象
func NewTaobaoAlitripBuyerGetRequest() *TaobaoAlitripBuyerGetAPIRequest {
	return &TaobaoAlitripBuyerGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripBuyerGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.buyer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripBuyerGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is RequestAxb Setter
// 敏感信息查询请求参数
func (r *TaobaoAlitripBuyerGetAPIRequest) SetRequestAxb(_requestAxb *RequestAxbDo) error {
	r._requestAxb = _requestAxb
	r.Set("request_axb", _requestAxb)
	return nil
}

// Get RequestAxb Getter
func (r TaobaoAlitripBuyerGetAPIRequest) GetRequestAxb() *RequestAxbDo {
	return r._requestAxb
}
