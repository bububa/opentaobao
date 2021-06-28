package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票代理商接口—确认出票是否成功 APIResponse
taobao.bus.agent.bookticket.confirm

代理商通过该接口通知汽车票系统订单出票结果。
*/
type TaobaoBusAgentBookticketConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusAgentBookticketConfirmResponse `json:"bus_agent_bookticket_confirm_response,omitempty"` 
    TaobaoBusAgentBookticketConfirmResponse
}

/* model for simplify = false
type TaobaoBusAgentBookticketConfirmResponse struct {

    // 是否确认成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

}
*/

type TaobaoBusAgentBookticketConfirmResponse struct {

    // 是否确认成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

}
