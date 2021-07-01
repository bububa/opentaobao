package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

/* TaobaoEticketMerchantMaDelay
凭证延期
taobao.eticket.merchant.ma.delay

订单延期 */
func TaobaoEticketMerchantMaDelay(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaDelayAPIRequest, session string) (*eticket.TaobaoEticketMerchantMaDelayAPIResponse, error) {
	var resp eticket.TaobaoEticketMerchantMaDelayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
