package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeRefuseVtwoAPIRequest
代理商拒绝改签v2--增加鉴权校验 API请求
taobao.train.agent.change.refuse.vtwo

代理商拒绝火车票改签服务 */
type TaobaoTrainAgentChangeRefuseVtwoAPIRequest struct {
	model.Params
	// 代理商拒绝改签参数
	_param *AgentRefuseChangeParam
}

// New
