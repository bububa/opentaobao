package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
综合交通多次退款接口 APIRequest
taobao.bus.agent.multiple.refund.confirm

此接口支持多次按照单客进行多次退款操作，只进行退款操作。
*/
type TaobaoBusAgentMultipleRefundConfirmRequest struct {
    model.Params

    // 入参
    paramAgentMultipleRefundRQ   *AgentMultipleRefundRq 

}

func NewTaobaoBusAgentMultipleRefundConfirmRequest() *TaobaoBusAgentMultipleRefundConfirmRequest{
    return &TaobaoBusAgentMultipleRefundConfirmRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoBusAgentMultipleRefundConfirmRequest) GetApiMethodName() string {
    return "taobao.bus.agent.multiple.refund.confirm"
}

func (r TaobaoBusAgentMultipleRefundConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoBusAgentMultipleRefundConfirmRequest) SetParamAgentMultipleRefundRQ(paramAgentMultipleRefundRQ *AgentMultipleRefundRq) error {
    r.paramAgentMultipleRefundRQ = paramAgentMultipleRefundRQ
    r.Set("param_agent_multiple_refund_r_q", paramAgentMultipleRefundRQ)
    return nil
}

func (r TaobaoBusAgentMultipleRefundConfirmRequest) GetParamAgentMultipleRefundRQ() *AgentMultipleRefundRq {
    return r.paramAgentMultipleRefundRQ
}

