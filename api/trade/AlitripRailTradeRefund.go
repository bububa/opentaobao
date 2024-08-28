package trade

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlitripRailTradeRefund 退票接口
// alitrip.rail.trade.refund
//
// 退票接口
func AlitripRailTradeRefund(ctx context.Context, clt *core.SDKClient, req *trade.AlitripRailTradeRefundAPIRequest, resp *trade.AlitripRailTradeRefundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
