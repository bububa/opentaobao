package opentrade

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest 高级定制商家传入改价信息 API请求
// taobao.miniapp.advanced.tradeinfo.price.modify
//
// 高级定制商家传入改价信息
type TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest struct {
	model.Params
	// 请求参数
	_req *SaveModifyPriceRequest
}

// NewTaobaoMiniappAdvancedTradeinfoPriceModifyRequest 初始化TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest对象
func NewTaobaoMiniappAdvancedTradeinfoPriceModifyRequest() *TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest {
	return &TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest) GetApiMethodName() string {
	return "taobao.miniapp.advanced.tradeinfo.price.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetReq is Req Setter
// 请求参数
func (r *TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest) SetReq(_req *SaveModifyPriceRequest) error {
	r._req = _req
	r.Set("req", _req)
	return nil
}

// GetReq Req Getter
func (r TaobaoMiniappAdvancedTradeinfoPriceModifyAPIRequest) GetReq() *SaveModifyPriceRequest {
	return r._req
}
