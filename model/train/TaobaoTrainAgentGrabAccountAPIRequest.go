package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentGrabAccountAPIRequest
代购抢代理商回传12306账号 API请求
taobao.train.agent.grab.account

火车票业务代购抢功能，代理商回传12306账号，用于自营抢票链路出票 */
type TaobaoTrainAgentGrabAccountAPIRequest struct {
	model.Params
	// 12306账户信息
	_accountParam *AccountParam
}

// New
