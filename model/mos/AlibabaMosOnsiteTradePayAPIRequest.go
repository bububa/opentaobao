package mos

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMosOnsiteTradePayAPIRequest
新商场当面付商户扫码付 API请求
alibaba.mos.onsite.trade.pay

收银员使用扫码设备读取用户“付款码”后，将二维码或条码信息通过本接口上送至喵街发起支付。 */
type AlibabaMosOnsiteTradePayAPIRequest struct {
	model.Params
	// 创建订单请求参数
	_onsiteTradePayRequest *OnsiteTradePayRequest
}

// NewAlibabaMosOnsiteTradePayRequest 初始化AlibabaMosOnsiteTradePayAPIRequest对象
func NewAlibabaMosOnsiteTradePayRequest() *AlibabaMosOnsiteTradePayAPIRequest {
	return &AlibabaMosOnsiteTradePayAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMosOnsiteTradePayAPIRequest) GetApiMethodName() string {
	return "alibaba.mos.onsite.trade.pay"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMosOnsiteTradePayAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OnsiteTradePayRequest Setter
// 创建订单请求参数
func (r *AlibabaMosOnsiteTradePayAPIRequest) SetOnsiteTradePayRequest(_onsiteTradePayRequest *OnsiteTradePayRequest) error {
	r._onsiteTradePayRequest = _onsiteTradePayRequest
	r.Set("onsite_trade_pay_request", _onsiteTradePayRequest)
	return nil
}

// Get OnsiteTradePayRequest Getter
func (r AlibabaMosOnsiteTradePayAPIRequest) GetOnsiteTradePayRequest() *OnsiteTradePayRequest {
	return r._onsiteTradePayRequest
}
