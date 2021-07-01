package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentTostationReceiveAPIRequest
线下票送票至车站代理商确认用户已取票服务 API请求
taobao.train.agent.tostation.receive

送票至车站的订单，代理商确认用户已取票 */
type TaobaoTrainAgentTostationReceiveAPIRequest struct {
	model.Params
	// 淘宝的主订单号
	_mainOrderId int64
	// 代理商id
	_agentId int64
}

// New
