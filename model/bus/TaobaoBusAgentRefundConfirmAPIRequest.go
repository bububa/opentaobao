package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退票和退款二合一接口 API请求
taobao.bus.agent.refund.confirm

1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
*/
type TaobaoBusAgentRefundConfirmAPIRequest struct {
    model.Params
    // 入参
    _paramAgentConfirmReturnAndRefundRQ   *AgentConfirmReturnAndRefundRq
}

// 初始化TaobaoBusAgentRefundConfirmAPIRequest对象
func NewTaobaoBusAgentRefundConfirmRequest() *TaobaoBusAgentRefundConfirmAPIRequest{
    return &TaobaoBusAgentRefundConfirmAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetApiMethodName() string {
    return "taobao.bus.agent.refund.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAgentConfirmReturnAndRefundRQ Setter
// 入参
func (r *TaobaoBusAgentRefundConfirmAPIRequest) SetParamAgentConfirmReturnAndRefundRQ(_paramAgentConfirmReturnAndRefundRQ *AgentConfirmReturnAndRefundRq) error {
    r._paramAgentConfirmReturnAndRefundRQ = _paramAgentConfirmReturnAndRefundRQ
    r.Set("param_agent_confirm_return_and_refund_r_q", _paramAgentConfirmReturnAndRefundRQ)
    return nil
}

// ParamAgentConfirmReturnAndRefundRQ Getter
func (r TaobaoBusAgentRefundConfirmAPIRequest) GetParamAgentConfirmReturnAndRefundRQ() *AgentConfirmReturnAndRefundRq {
    return r._paramAgentConfirmReturnAndRefundRQ
}
