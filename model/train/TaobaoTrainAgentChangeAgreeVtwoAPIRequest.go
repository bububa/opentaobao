package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeAgreeVtwoAPIRequest
代理商同意改签v2--增加鉴权校验 API请求
taobao.train.agent.change.agree.vtwo

代理商同意改签接口服务 */
type TaobaoTrainAgentChangeAgreeVtwoAPIRequest struct {
	model.Params
	// 代理商同意改签参数
	_param *AgentAgreeChangeParam
}

// New
