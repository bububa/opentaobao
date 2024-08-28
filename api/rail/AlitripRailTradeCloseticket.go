package rail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailTradeCloseticket 出票失败关单接口
// alitrip.rail.trade.closeticket
//
// 出票成功回调接口
func AlitripRailTradeCloseticket(ctx context.Context, clt *core.SDKClient, req *rail.AlitripRailTradeCloseticketAPIRequest, resp *rail.AlitripRailTradeCloseticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
