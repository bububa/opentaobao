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

// New
