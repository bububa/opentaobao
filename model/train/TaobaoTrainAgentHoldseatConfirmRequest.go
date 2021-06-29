package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认占座是否成功 API请求
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功
*/
type TaobaoTrainAgentHoldseatConfirmRequest struct {
    model.Params
    // 占座入参
    _holdSeatParam   *HoldSeatParam
}

// 初始化TaobaoTrainAgentHoldseatConfirmRequest对象
func NewTaobaoTrainAgentHoldseatConfirmRequest() *TaobaoTrainAgentHoldseatConfirmRequest{
    return &TaobaoTrainAgentHoldseatConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoTrainAgentHoldseatConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.holdseat.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoTrainAgentHoldseatConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// HoldSeatParam Setter
// 占座入参
func (r *TaobaoTrainAgentHoldseatConfirmRequest) SetHoldSeatParam(_holdSeatParam *HoldSeatParam) error {
    r._holdSeatParam = _holdSeatParam
    r.Set("hold_seat_param", _holdSeatParam)
    return nil
}

// HoldSeatParam Getter
func (r TaobaoTrainAgentHoldseatConfirmRequest) GetHoldSeatParam() *HoldSeatParam {
    return r._holdSeatParam
}
