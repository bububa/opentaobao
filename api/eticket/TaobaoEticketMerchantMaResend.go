package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoEticketMerchantMaResend 电子凭证重发回调接口
// taobao.eticket.merchant.ma.resend
//
// 码商重发电子凭证回调接口
func TaobaoEticketMerchantMaResend(clt *core.SDKClient, req *eticket.TaobaoEticketMerchantMaResendAPIRequest, resp *eticket.TaobaoEticketMerchantMaResendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
