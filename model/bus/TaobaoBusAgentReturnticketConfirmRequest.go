package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调退票 APIRequest
taobao.bus.agent.returnticket.confirm

商家通过TOP接口调用来回传退票状态
*/
type TaobaoBusAgentReturnticketConfirmRequest struct {
    model.Params

    // 退票入参
    paramAgentConfirmReturnRQ   *AgentConfirmReturnRq 

}

func NewTaobaoBusAgentReturnticketConfirmRequest() *TaobaoBusAgentReturnticketConfirmRequest{
    return &TaobaoBusAgentReturnticketConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusAgentReturnticketConfirmRequest) GetApiMethodName() string {
    return "taobao.bus.agent.returnticket.confirm"
}

func (r TaobaoBusAgentReturnticketConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusAgentReturnticketConfirmRequest) SetParamAgentConfirmReturnRQ(paramAgentConfirmReturnRQ *AgentConfirmReturnRq) error {
    r.paramAgentConfirmReturnRQ = paramAgentConfirmReturnRQ
    r.Set("param_agent_confirm_return_r_q", paramAgentConfirmReturnRQ)
    return nil
}

func (r TaobaoBusAgentReturnticketConfirmRequest) GetParamAgentConfirmReturnRQ() *AgentConfirmReturnRq {
    return r.paramAgentConfirmReturnRQ
}

