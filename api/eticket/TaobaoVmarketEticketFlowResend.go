package eticket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketFlowResend 业务重新触发发码短信
// taobao.vmarket.eticket.flow.resend
//
// 业务重新触发发码短信
func TaobaoVmarketEticketFlowResend(clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFlowResendAPIRequest, resp *eticket.TaobaoVmarketEticketFlowResendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
