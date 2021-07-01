package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeGetVtwoAPIRequest
获取改签单详情v2--增加鉴权校验 API请求
taobao.train.agent.change.get.vtwo

卖家获取待处理的改签单详情 */
type TaobaoTrainAgentChangeGetVtwoAPIRequest struct {
	model.Params
	// 代理商id
	_agentId int64
	// 申请单id
	_applyId int64
}

// New
