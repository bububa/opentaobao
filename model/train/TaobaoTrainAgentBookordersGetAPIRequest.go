package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentBookordersGetAPIRequest
代理商获取待出票订单列表 API请求
taobao.train.agent.bookorders.get

代理商获取待出票订单列表，只返回订单号 */
type TaobaoTrainAgentBookordersGetAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
}

// New
