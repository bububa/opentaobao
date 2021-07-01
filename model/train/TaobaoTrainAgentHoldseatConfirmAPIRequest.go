package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentHoldseatConfirmAPIRequest
火车票代理商接口——确认占座是否成功 API请求
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功 */
type TaobaoTrainAgentHoldseatConfirmAPIRequest struct {
	model.Params
	// 占座入参
	_holdSeatParam *HoldSeatParam
}

// New
