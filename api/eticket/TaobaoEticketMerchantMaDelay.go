package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaDelay 凭证延期
// taobao.eticket.merchant.ma.delay
//
// 订单延期
func TaobaoEticketMerchantMaDelay(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaDelayAPIRequest, resp *eticket.TaobaoEticketMerchantMaDelayAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
