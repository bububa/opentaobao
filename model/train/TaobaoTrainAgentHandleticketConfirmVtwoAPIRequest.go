package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest
代理商出票中v2--增加鉴权校验 API请求
taobao.train.agent.handleticket.confirm.vtwo

代理商出票中 */
type TaobaoTrainAgentHandleticketConfirmVtwoAPIRequest struct {
	model.Params
	// 扩展参数
	_extendParams string
	// 主站id
	_mainOrderId int64
	// 代理商id
	_sellerId int64
}

// New
