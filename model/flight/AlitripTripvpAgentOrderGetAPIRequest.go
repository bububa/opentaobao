package flight

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlitripTripvpAgentOrderGetAPIRequest
廉航辅营正向订单查询详情接口 API请求
alitrip.tripvp.agent.order.get

【国际机票】查询辅营订单详情 */
type AlitripTripvpAgentOrderGetAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 辅营的订单号
	_tradeOrderId int64
}

// New
