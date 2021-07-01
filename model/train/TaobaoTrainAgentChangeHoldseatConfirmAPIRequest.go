package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTrainAgentChangeHoldseatConfirmAPIRequest
火车票代理商接口——确认改签占座是否成功 API请求
taobao.train.agent.change.holdseat.confirm

火车票代理商接口——确认改签占座是否成功 */
type TaobaoTrainAgentChangeHoldseatConfirmAPIRequest struct {
	model.Params
	// 改签占座入参
	_changeHoldSeatParam *ChangeHoldSeatParam
}

// New
