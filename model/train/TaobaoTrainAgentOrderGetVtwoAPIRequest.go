package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentOrderGetVtwoAPIRequest
代理商获取订单信息回调APIv2--增加鉴权校验 API请求
taobao.train.agent.order.get.vtwo

代理商获取订单信息回调API */
type TaobaoTrainAgentOrderGetVtwoAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// New
