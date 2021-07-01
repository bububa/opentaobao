package ieagency

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripIeAgentRefundNewReceiveAPIRequest
商家退票受理申请(对外) API请求
taobao.alitrip.ie.agent.refund.new.receive

允许代理商通过top接口受理退票申请 */
type TaobaoAlitripIeAgentRefundNewReceiveAPIRequest struct {
	model.Params
	// 退票申请单id
	_applyId int64
	// 代理商id
	_agentId int64
	// 订单id
	_orderId int64
}

// New
