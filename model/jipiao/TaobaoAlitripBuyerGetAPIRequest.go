package jipiao

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripBuyerGetAPIRequest 敏感信息查询 API请求
// taobao.alitrip.buyer.get
//
// 针对商家提供统一的TOP接口，可以根据订单获取订单对应买家联系电话（阿里小号）。
type TaobaoAlitripBuyerGetAPIRequest struct {
	model.Params
	// 敏感信息查询请求参数
	_requestAxb *RequestAxbDo
}

// NewTaobaoAlitripBuyerGetRequest 初始化TaobaoAlitripBuyerGetAPIRequest对象
func NewTaobaoAlitripBuyerGetRequest() *TaobaoAlitripBuyerGetAPIRequest {
	return &TaobaoAlitripBuyerGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoAlitripBuyerGetAPIRequest) Reset() {
	r._requestAxb = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoAlitripBuyerGetAPIRequest) GetApiMethodName() string {
	return "taobao.alitrip.buyer.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoAlitripBuyerGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoAlitripBuyerGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetRequestAxb is RequestAxb Setter
// 敏感信息查询请求参数
func (r *TaobaoAlitripBuyerGetAPIRequest) SetRequestAxb(_requestAxb *RequestAxbDo) error {
	r._requestAxb = _requestAxb
	r.Set("request_axb", _requestAxb)
	return nil
}

// GetRequestAxb RequestAxb Getter
func (r TaobaoAlitripBuyerGetAPIRequest) GetRequestAxb() *RequestAxbDo {
	return r._requestAxb
}

var poolTaobaoAlitripBuyerGetAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoAlitripBuyerGetRequest()
	},
}

// GetTaobaoAlitripBuyerGetRequest 从 sync.Pool 获取 TaobaoAlitripBuyerGetAPIRequest
func GetTaobaoAlitripBuyerGetAPIRequest() *TaobaoAlitripBuyerGetAPIRequest {
	return poolTaobaoAlitripBuyerGetAPIRequest.Get().(*TaobaoAlitripBuyerGetAPIRequest)
}

// ReleaseTaobaoAlitripBuyerGetAPIRequest 将 TaobaoAlitripBuyerGetAPIRequest 放入 sync.Pool
func ReleaseTaobaoAlitripBuyerGetAPIRequest(v *TaobaoAlitripBuyerGetAPIRequest) {
	v.Reset()
	poolTaobaoAlitripBuyerGetAPIRequest.Put(v)
}
