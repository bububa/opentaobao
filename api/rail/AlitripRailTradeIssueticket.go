package rail

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/rail"
)

// AlitripRailTradeIssueticket 德铁出票成功接口
// alitrip.rail.trade.issueticket
//
// 出票成功回调接口
func AlitripRailTradeIssueticket(clt *core.SDKClient, req *rail.AlitripRailTradeIssueticketAPIRequest, resp *rail.AlitripRailTradeIssueticketAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
