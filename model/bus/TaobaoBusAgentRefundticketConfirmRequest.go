package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家top回调退款明细 API请求
taobao.bus.agent.refundticket.confirm

商家通过top回调告知平台退款明细
*/
type TaobaoBusAgentRefundticketConfirmRequest struct {
    model.Params
    // 退款入参
    _paramAgentConfirmRefundRQ   *AgentConfirmRefundRq
}

// 初始化TaobaoBusAgentRefundticketConfirmRequest对象
func NewTaobaoBusAgentRefundticketConfirmRequest() *TaobaoBusAgentRefundticketConfirmRequest{
    return &TaobaoBusAgentRefundticketConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentRefundticketConfirmRequest) GetApiMethodName() string {
    return "taobao.bus.agent.refundticket.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentRefundticketConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAgentConfirmRefundRQ Setter
// 退款入参
func (r *TaobaoBusAgentRefundticketConfirmRequest) SetParamAgentConfirmRefundRQ(_paramAgentConfirmRefundRQ *AgentConfirmRefundRq) error {
    r._paramAgentConfirmRefundRQ = _paramAgentConfirmRefundRQ
    r.Set("param_agent_confirm_refund_r_q", _paramAgentConfirmRefundRQ)
    return nil
}

// ParamAgentConfirmRefundRQ Getter
func (r TaobaoBusAgentRefundticketConfirmRequest) GetParamAgentConfirmRefundRQ() *AgentConfirmRefundRq {
    return r._paramAgentConfirmRefundRQ
}
