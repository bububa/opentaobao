package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退票和退款二合一接口 APIRequest
taobao.bus.agent.refund.confirm

1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
*/
type TaobaoBusAgentRefundConfirmRequest struct {
    model.Params

    // 入参
    paramAgentConfirmReturnAndRefundRQ   *AgentConfirmReturnAndRefundRq 

}

func NewTaobaoBusAgentRefundConfirmRequest() *TaobaoBusAgentRefundConfirmRequest{
    return &TaobaoBusAgentRefundConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusAgentRefundConfirmRequest) GetApiMethodName() string {
    return "taobao.bus.agent.refund.confirm"
}

func (r TaobaoBusAgentRefundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusAgentRefundConfirmRequest) SetParamAgentConfirmReturnAndRefundRQ(paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq) error {
    r.paramAgentConfirmReturnAndRefundRQ = paramAgentConfirmReturnAndRefundRQ
    r.Set("param_agent_confirm_return_and_refund_r_q", paramAgentConfirmReturnAndRefundRQ)
    return nil
}

func (r TaobaoBusAgentRefundConfirmRequest) GetParamAgentConfirmReturnAndRefundRQ() *AgentConfirmReturnAndRefundRq {
    return r.paramAgentConfirmReturnAndRefundRQ
}

