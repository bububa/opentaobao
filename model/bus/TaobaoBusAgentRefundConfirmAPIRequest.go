package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusAgentRefundConfirmAPIRequest
汽车票退票和退款二合一接口 API请求
taobao.bus.agent.refund.confirm

1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。 */
type TaobaoBusAgentRefundConfirmAPIRequest struct {
	model.Params
	// 入参
	_paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq
}

// New
