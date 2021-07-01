package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentBookordersGetVtwoAPIRequest
代理商获取待出票订单列表v2--增加鉴权校验 API请求
taobao.train.agent.bookorders.get.vtwo

代理商获取待出票订单列表，只返回订单号 */
type TaobaoTrainAgentBookordersGetVtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
}

// New
