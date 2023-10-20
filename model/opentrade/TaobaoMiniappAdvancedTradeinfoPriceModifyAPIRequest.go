package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaominiappadvancedtradeinfopricemodifyAPIRequest 高级定制商家传入改价信息 API请求
// taobao.miniapp.advanced.tradeinfo.price.modify
//
// 高级定制商家传入改价信息
type TaobaominiappadvancedtradeinfopricemodifyAPIRequest struct {
	model.Params
	// 请求参数
	_req *SaveModifyPriceRequest
}

// NewTaobaominiappadvancedtradeinfopricemodifyRequest 初始化TaobaominiappadvancedtradeinfopricemodifyAPIRequest对象
func NewTaobaominiappadvancedtradeinfopricemodifyRequest() *TaobaominiappadvancedtradeinfopricemodifyAPIRequest {
	return &TaobaominiappadvancedtradeinfopricemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaominiappadvancedtradeinfopricemodifyAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.advanced.tradeinfo.price.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaominiappadvancedtradeinfopricemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaominiappadvancedtradeinfopricemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求参数
func (r *TaobaominiappadvancedtradeinfopricemodifyAPIRequest) SetReq(_req *SaveModifyPriceRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaominiappadvancedtradeinfopricemodifyAPIRequest) GetReq() *SaveModifyPriceRequest {
	return r._req
}
