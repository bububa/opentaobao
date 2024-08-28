package eticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eticket"
)

// TaobaoVmarketEticketFlowConsume 无交易类凭证核销
// taobao.vmarket.eticket.flow.consume
//
// 无交易类凭证核销
func TaobaoVmarketEticketFlowConsume(ctx context.Context, clt *core.SDKClient, req *eticket.TaobaoVmarketEticketFlowConsumeAPIRequest, resp *eticket.TaobaoVmarketEticketFlowConsumeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
