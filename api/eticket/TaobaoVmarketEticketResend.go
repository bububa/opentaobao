package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketResend 外部合作商家重发电子凭证回调接口
// taobao.vmarket.eticket.resend
//
// 外部合作商家重发电子凭证回调接口
func TaobaoVmarketEticketResend(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketResendAPIRequest, resp *eticket.TaobaoVmarketEticketResendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
