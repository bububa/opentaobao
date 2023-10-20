package tmallcar

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallAliautoFulfillmentAuthCheckAPIRequest 商家鉴权 API请求
// tmall.aliauto.fulfillment.auth.check
//
// 商家鉴权
type TmallAliautoFulfillmentAuthCheckAPIRequest struct {
	model.Params
	// 请求参数
	_req *AuthCheckReq
}

// NewTmallAliautoFulfillmentAuthCheckRequest 初始化TmallAliautoFulfillmentAuthCheckAPIRequest对象
func NewTmallAliautoFulfillmentAuthCheckRequest() *TmallAliautoFulfillmentAuthCheckAPIRequest {
	return &TmallAliautoFulfillmentAuthCheckAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallAliautoFulfillmentAuthCheckAPIRequest) GetApiMethodName() string {
	return "tmall.aliauto.fulfillment.auth.check"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallAliautoFulfillmentAuthCheckAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallAliautoFulfillmentAuthCheckAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求参数
func (r *TmallAliautoFulfillmentAuthCheckAPIRequest) SetReq(_req *AuthCheckReq) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TmallAliautoFulfillmentAuthCheckAPIRequest) GetReq() *AuthCheckReq {
	return r._req
}
