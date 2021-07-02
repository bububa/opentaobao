package train

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTrainAgentHoldseatConfirmAPIRequest 火车票代理商接口——确认占座是否成功 API请求
// taobao.train.agent.holdseat.confirm
//
// 火车票代理商接口——确认占座是否成功
type TaobaoTrainAgentHoldseatConfirmAPIRequest struct {
	model.Params
	// 占座入参
	_holdSeatParam *HoldSeatParam
}

// NewTaobaoTrainAgentHoldseatConfirmRequest 初始化TaobaoTrainAgentHoldseatConfirmAPIRequest对象
func NewTaobaoTrainAgentHoldseatConfirmRequest() *TaobaoTrainAgentHoldseatConfirmAPIRequest {
	return &TaobaoTrainAgentHoldseatConfirmAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHoldseatConfirmAPIRequest) GetApiMethodName() string {
	return "taobao.train.agent.holdseat.confirm"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHoldseatConfirmAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is HoldSeatParam Setter
// 占座入参
func (r *TaobaoTrainAgentHoldseatConfirmAPIRequest) SetHoldSeatParam(_holdSeatParam *HoldSeatParam) error {
	r._holdSeatParam = _holdSeatParam
	r.Set("hold_seat_param", _holdSeatParam)
	return nil
}

// Get HoldSeatParam Getter
func (r TaobaoTrainAgentHoldseatConfirmAPIRequest) GetHoldSeatParam() *HoldSeatParam {
	return r._holdSeatParam
}
