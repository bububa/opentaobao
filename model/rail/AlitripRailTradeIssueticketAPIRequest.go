package rail

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripRailTradeIssueticketAPIRequest
德铁出票成功接口 API请求
alitrip.rail.trade.issueticket

出票成功回调接口 */
type AlitripRailTradeIssueticketAPIRequest struct {
	model.Params
	// 代理商订单号
	_agentOrderId string
	// 平台订单号
	_tpOrderId int64
	// 代理商id
	_agentId int64
	// pnr票号有则填，无则空
	_ticketNo string
}

// New
