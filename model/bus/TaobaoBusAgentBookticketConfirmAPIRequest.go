package bus

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBusAgentBookticketConfirmAPIRequest
汽车票代理商接口—确认出票是否成功 API请求
taobao.bus.agent.bookticket.confirm

代理商通过该接口通知汽车票系统订单出票结果。 */
type TaobaoBusAgentBookticketConfirmAPIRequest struct {
	model.Params
	// 请求对象
	_paramAgentConfirmBookRQ *AgentConfirmBookRq
}

// New
