package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentGetRefundAPIRequest
代理商获取订单退票信息 API请求
taobao.train.agent.get.refund

代理商获取订单信息回调API */
type TaobaoTrainAgentGetRefundAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// New
