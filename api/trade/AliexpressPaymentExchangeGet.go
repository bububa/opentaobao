package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AliexpressPaymentExchangeGet getExchange
// aliexpress.payment.exchange.get
//
// 提供国际汇率服务
func AliexpressPaymentExchangeGet(clt *core.SDKClient, req *trade.AliexpressPaymentExchangeGetAPIRequest, resp *trade.AliexpressPaymentExchangeGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
