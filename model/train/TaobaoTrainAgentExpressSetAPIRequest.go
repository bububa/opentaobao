package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentExpressSetAPIRequest
线下票回填物流信息 API请求
taobao.train.agent.express.set

线下票回填物流信息服务 */
type TaobaoTrainAgentExpressSetAPIRequest struct {
	model.Params
	// 订单号
	_mainOrderId int64
	// 物流单号
	_expressId string
	// 发货地址
	_addr string
	// 手机号
	_mobile string
	// 代理商id
	_agentId int64
	// 物流公司:SF,EMS
	_expressName string
}

// New
