package tbtrade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoTradeConfirmfeeGet 获取交易确认收货费用
// taobao.trade.confirmfee.get
//
// 获取交易确认收货费用，可以获取主订单或子订单的确认收货费用
func TaobaoTradeConfirmfeeGet(ctx context.Context, clt *core.SDKClient, req *tbtrade.TaobaoTradeConfirmfeeGetAPIRequest, resp *tbtrade.TaobaoTradeConfirmfeeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
