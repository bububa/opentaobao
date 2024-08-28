package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketFlowResend 业务重新触发发码短信
// taobao.vmarket.eticket.flow.resend
//
// 业务重新触发发码短信
func TaobaoVmarketEticketFlowResend(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFlowResendAPIRequest, resp *eticket.TaobaoVmarketEticketFlowResendAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
