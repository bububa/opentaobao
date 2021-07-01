package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentOrderPayAPIRequest
代购订单代付接口 API请求
taobao.train.agent.order.pay

代购订单代付接口 */
type TaobaoTrainAgentOrderPayAPIRequest struct {
	model.Params
	// 入参对象
	_agentPayOrderParam *AgentPayOrderParam
}

// New
