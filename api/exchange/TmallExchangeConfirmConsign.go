package exchange

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeConfirmConsign 换货商家确认收货并发货
// tmall.exchange.confirm.consign
//
// 卖家确认收货并发货
func TmallExchangeConfirmConsign(ctx context.Context, clt *core.SDKClient, req *exchange.TmallExchangeConfirmConsignAPIRequest, resp *exchange.TmallExchangeConfirmConsignAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
