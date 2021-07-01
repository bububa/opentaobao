package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentReturnordersGetVtwoAPIRequest
获取待退票的订单v2--增加鉴权校验 API请求
taobao.train.agent.returnorders.get.vtwo

代理商用来获取待退票的订单列表及数量，防止代理商掉单。 */
type TaobaoTrainAgentReturnordersGetVtwoAPIRequest struct {
	model.Params
	// 卖家ID
	_agentId int64
	// 0 线上退票 1线下退票
	_offline int64
}

// New
