package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentOrderGetAPIRequest
代理商获取订单信息回调API API请求
taobao.train.agent.order.get

代理商获取订单信息回调API */
type TaobaoTrainAgentOrderGetAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// New
