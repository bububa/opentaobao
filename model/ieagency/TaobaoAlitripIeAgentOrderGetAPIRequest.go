package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentOrderGetAPIRequest
【国际机票】查询订单详情 API请求
taobao.alitrip.ie.agent.order.get

根据订单ID查询订单详情 */
type TaobaoAlitripIeAgentOrderGetAPIRequest struct {
	model.Params
	// 代理商ID
	_agentId int64
	// 交易订单ID
	_tradeOrderId int64
}

// New
