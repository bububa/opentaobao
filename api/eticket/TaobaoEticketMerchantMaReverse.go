package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaReverse 电子凭证冲正接口
// taobao.eticket.merchant.ma.reverse
//
// 电子凭证平台冲正接口
func TaobaoEticketMerchantMaReverse(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaReverseAPIRequest, resp *eticket.TaobaoEticketMerchantMaReverseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
