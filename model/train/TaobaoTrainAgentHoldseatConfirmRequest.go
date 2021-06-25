package train

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
火车票代理商接口——确认占座是否成功 APIRequest
taobao.train.agent.holdseat.confirm

火车票代理商接口——确认占座是否成功
*/
type TaobaoTrainAgentHoldseatConfirmRequest struct {
    model.Params

    // 占座入参
    holdSeatParam   *HoldSeatParam 

}

func NewTaobaoTrainAgentHoldseatConfirmRequest() *TaobaoTrainAgentHoldseatConfirmRequest{
    return &TaobaoTrainAgentHoldseatConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoTrainAgentHoldseatConfirmRequest) GetApiMethodName() string {
    return "taobao.train.agent.holdseat.confirm"
}

func (r TaobaoTrainAgentHoldseatConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoTrainAgentHoldseatConfirmRequest) SetHoldSeatParam(holdSeatParam *HoldSeatParam) error {
    r.holdSeatParam = holdSeatParam
    r.Set("hold_seat_param", holdSeatParam)
    return nil
}

func (r TaobaoTrainAgentHoldseatConfirmRequest) GetHoldSeatParam() *HoldSeatParam {
    return r.holdSeatParam
}

