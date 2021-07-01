package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTripvpAgentOrderIssueAPIRequest
廉航辅营正向订单出货接口 API请求
alitrip.tripvp.agent.order.issue

廉航辅营正向订单出货接口 */
type AlitripTripvpAgentOrderIssueAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 出货通知信息
	_issueProOrderVo *IssueProOrderVo
}

// New
