package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentDirectCompensateAPIRequest
火车票代理商接口——订单关闭实际出票成功审计接口 API请求
taobao.train.agent.direct.compensate

代购直连订单平台关单但是代理商出票成功补偿接口 */
type TaobaoTrainAgentDirectCompensateAPIRequest struct {
	model.Params
	// 出票成功补偿入参
	_compensateParam *CompensateParam
}

// New
