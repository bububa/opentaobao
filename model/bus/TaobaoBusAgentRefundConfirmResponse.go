package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
汽车票退票和退款二合一接口 APIResponse
taobao.bus.agent.refund.confirm

1.商家退票成功后，回调飞猪平台汽车票退票接口，平台进行退票和退款操作。
*/
type TaobaoBusAgentRefundConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusAgentRefundConfirmResponse `json:"bus_agent_refund_confirm_response,omitempty"` 
    TaobaoBusAgentRefundConfirmResponse
}

/* model for simplify = false
type TaobaoBusAgentRefundConfirmResponse struct {

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 错误描述
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 退票回调是否收到
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusAgentRefundConfirmResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 错误描述
    ResultMsg   string `json:"result_msg,omitempty"`

    // 退票回调是否收到
    IsSuccess   bool `json:"is_success,omitempty"`

}
