package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaAvailable 电子凭证核销前校验接口
// taobao.eticket.merchant.ma.available
//
// 商家验码之前的调用接口，用来判断是否可以进行核销操作
func TaobaoEticketMerchantMaAvailable(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaAvailableAPIRequest, resp *eticket.TaobaoEticketMerchantMaAvailableAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
