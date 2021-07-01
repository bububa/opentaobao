package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeordersGetVtwoAPIRequest
获取待改签订单v2--增加鉴权校验 API请求
taobao.train.agent.changeorders.get.vtwo

代理商用来获取待改签的订单列表及数量，防止代理商掉单。 */
type TaobaoTrainAgentChangeordersGetVtwoAPIRequest struct {
	model.Params
	// 卖家id
	_agentId int64
}

// New
