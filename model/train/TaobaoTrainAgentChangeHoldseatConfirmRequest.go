package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认改签占座是否成功 API请求
taobao.train.agent.change.holdseat.confirm

火车票代理商接口——确认改签占座是否成功
*/
type TaobaoTrainAgentChangeHoldseatConfirmAPIRequest struct {
    model.Params
    // 改签占座入参
    _changeHoldSeatParam   *ChangeHoldSeatParam
}

// 初始化TaobaoTrainAgentChangeHoldseatConfirmAPIRequest对象
func NewTaobaoTrainAgentChangeHoldseatConfirmRequest() *TaobaoTrainAgentChangeHoldseatConfirmAPIRequest{
    return &TaobaoTrainAgentChangeHoldseatConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentChangeHoldseatConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.holdseat.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentChangeHoldseatConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ChangeHoldSeatParam Setter
// 改签占座入参
func (r *TaobaoTrainAgentChangeHoldseatConfirmAPIRequest) SetChangeHoldSeatParam(_changeHoldSeatParam *ChangeHoldSeatParam) error {
    r._changeHoldSeatParam = _changeHoldSeatParam
    r.Set("change_hold_seat_param", _changeHoldSeatParam)
    return nil
}

// ChangeHoldSeatParam Getter
func (r TaobaoTrainAgentChangeHoldseatConfirmAPIRequest) GetChangeHoldSeatParam() *ChangeHoldSeatParam {
    return r._changeHoldSeatParam
}
