package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
商家回调退票 APIResponse
taobao.bus.agent.returnticket.confirm

商家通过TOP接口调用来回传退票状态
*/
type TaobaoBusAgentReturnticketConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusAgentReturnticketConfirmResponse `json:"bus_agent_returnticket_confirm_response,omitempty"` 
    TaobaoBusAgentReturnticketConfirmResponse
}

/* model for simplify = false
type TaobaoBusAgentReturnticketConfirmResponse struct {

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 是否确认成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusAgentReturnticketConfirmResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 是否确认成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
