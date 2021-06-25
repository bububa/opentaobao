package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认改签占座是否成功 APIRequest
taobao.train.agent.change.holdseat.confirm

火车票代理商接口——确认改签占座是否成功
*/
type TaobaoTrainAgentChangeHoldseatConfirmRequest struct {
    model.Params

    // 改签占座入参
    changeHoldSeatParam   *ChangeHoldSeatParam 

}

func NewTaobaoTrainAgentChangeHoldseatConfirmRequest() *TaobaoTrainAgentChangeHoldseatConfirmRequest{
    return &TaobaoTrainAgentChangeHoldseatConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentChangeHoldseatConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.change.holdseat.confirm"
}

func (r TaobaoTrainAgentChangeHoldseatConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentChangeHoldseatConfirmRequest) SetChangeHoldSeatParam(changeHoldSeatParam *ChangeHoldSeatParam) error {
    r.changeHoldSeatParam = changeHoldSeatParam
    r.Set("change_hold_seat_param", changeHoldSeatParam)
    return nil
}

func (r TaobaoTrainAgentChangeHoldseatConfirmRequest) GetChangeHoldSeatParam() *ChangeHoldSeatParam {
    return r.changeHoldSeatParam
}

