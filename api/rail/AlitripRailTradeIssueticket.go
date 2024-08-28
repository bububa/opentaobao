package rail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailTradeIssueticket 德铁出票成功接口
// alitrip.rail.trade.issueticket
//
// 出票成功回调接口
func AlitripRailTradeIssueticket(ctx context.Context, clt *core.SDKClient, req *rail.AlitripRailTradeIssueticketAPIRequest, resp *rail.AlitripRailTradeIssueticketAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
