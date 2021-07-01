package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

/* AliexpressPaymentExchangeGet
getExchange
aliexpress.payment.exchange.get

提供国际汇率服务 */
func AliexpressPaymentExchangeGet(clt *core.SDKClient, req *trade.AliexpressPaymentExchangeGetAPIRequest, session string) (*trade.AliexpressPaymentExchangeGetAPIResponse, error) {
	var resp trade.AliexpressPaymentExchangeGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
