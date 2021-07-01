package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHandleticketConfirmAPIRequest
代理商出票中 API请求
taobao.train.agent.handleticket.confirm

代理商出票中 */
type TaobaoTrainAgentHandleticketConfirmAPIRequest struct {
	model.Params
	// 扩展参数
	_extendParams string
	// 主站id
	_mainOrderId int64
	// 代理商id
	_sellerId int64
}

// New
