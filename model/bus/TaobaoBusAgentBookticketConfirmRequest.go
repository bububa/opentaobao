package bus

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
汽车票代理商接口—确认出票是否成功 API请求
taobao.bus.agent.bookticket.confirm

代理商通过该接口通知汽车票系统订单出票结果。
*/
type TaobaoBusAgentBookticketConfirmRequest struct {
    model.Params
    // 请求对象
    paramAgentConfirmBookRQ   *AgentConfirmBookRq
}

// 初始化TaobaoBusAgentBookticketConfirmRequest对象
func NewTaobaoBusAgentBookticketConfirmRequest() *TaobaoBusAgentBookticketConfirmRequest{
    return &TaobaoBusAgentBookticketConfirmRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoBusAgentBookticketConfirmRequest) GetApiMethodName() string {
    return "taobao.bus.agent.bookticket.confirm"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoBusAgentBookticketConfirmRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamAgentConfirmBookRQ Setter
// 请求对象
func (r *TaobaoBusAgentBookticketConfirmRequest) SetParamAgentConfirmBookRQ(paramAgentConfirmBookRQ *AgentConfirmBookRq) error {
    r.paramAgentConfirmBookRQ = paramAgentConfirmBookRQ
    r.Set("param_agent_confirm_book_r_q", paramAgentConfirmBookRQ)
    return nil
}

// ParamAgentConfirmBookRQ Getter
func (r TaobaoBusAgentBookticketConfirmRequest) GetParamAgentConfirmBookRQ() *AgentConfirmBookRq {
    return r.paramAgentConfirmBookRQ
}
