package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusAgentReturnticketConfirmAPIRequest
商家回调退票 API请求
taobao.bus.agent.returnticket.confirm

商家通过TOP接口调用来回传退票状态 */
type TaobaoBusAgentReturnticketConfirmAPIRequest struct {
	model.Params
	// 退票入参
	_paramAgentConfirmReturnRQ *AgentConfirmReturnRq
}

// New
