package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentTostationConfirmAPIRequest
线下票确认送票至车站服务 API请求
taobao.train.agent.tostation.confirm

送票至车站的订单，代理商确认配送到站 */
type TaobaoTrainAgentTostationConfirmAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// New
