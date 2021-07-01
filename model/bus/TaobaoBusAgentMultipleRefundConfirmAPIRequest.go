package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusAgentMultipleRefundConfirmAPIRequest
综合交通多次退款接口 API请求
taobao.bus.agent.multiple.refund.confirm

此接口支持多次按照单客进行多次退款操作，只进行退款操作。 */
type TaobaoBusAgentMultipleRefundConfirmAPIRequest struct {
	model.Params
	// 入参
	_paramAgentMultipleRefundRQ *AgentMultipleRefundRq
}

// New
