package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeordersGetAPIRequest
获取待改签订单 API请求
taobao.train.agent.changeorders.get

代理商用来获取待改签的订单列表及数量，防止代理商掉单。 */
type TaobaoTrainAgentChangeordersGetAPIRequest struct {
	model.Params
	// 卖家id
	_agentId int64
}

// New
