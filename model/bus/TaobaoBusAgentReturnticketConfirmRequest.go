package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家回调退票 API请求
taobao.bus.agent.returnticket.confirm

商家通过TOP接口调用来回传退票状态
*/
type TaobaoBusAgentReturnticketConfirmAPIRequest struct {
    model.Params
    // 退票入参
    _paramAgentConfirmReturnRQ   *AgentConfirmReturnRq
}

// 初始化TaobaoBusAgentReturnticketConfirmAPIRequest对象
func NewTaobaoBusAgentReturnticketConfirmRequest() *TaobaoBusAgentReturnticketConfirmAPIRequest{
    return &TaobaoBusAgentReturnticketConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.bus.agent.returnticket.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAgentConfirmReturnRQ Setter
// 退票入参
func (r *TaobaoBusAgentReturnticketConfirmAPIRequest) SetParamAgentConfirmReturnRQ(_paramAgentConfirmReturnRQ *AgentConfirmReturnRq) error {
    r._paramAgentConfirmReturnRQ = _paramAgentConfirmReturnRQ
    r.Set("param_agent_confirm_return_r_q", _paramAgentConfirmReturnRQ)
    return nil
}

// ParamAgentConfirmReturnRQ Getter
func (r TaobaoBusAgentReturnticketConfirmAPIRequest) GetParamAgentConfirmReturnRQ() *AgentConfirmReturnRq {
    return r._paramAgentConfirmReturnRQ
}
