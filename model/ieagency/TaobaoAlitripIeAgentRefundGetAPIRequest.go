package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundGetAPIRequest
获取退票申请详情 API请求
taobao.alitrip.ie.agent.refund.get

获取退票申请详情 */
type TaobaoAlitripIeAgentRefundGetAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
}

// New
