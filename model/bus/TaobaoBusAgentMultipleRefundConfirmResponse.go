package bus

import (
    "github.com/bububa/opentaobao/model"
)

/* 
综合交通多次退款接口 APIResponse
taobao.bus.agent.multiple.refund.confirm

此接口支持多次按照单客进行多次退款操作，只进行退款操作。
*/
type TaobaoBusAgentMultipleRefundConfirmAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoBusAgentMultipleRefundConfirmResponse `json:"bus_agent_multiple_refund_confirm_response,omitempty"` 
    TaobaoBusAgentMultipleRefundConfirmResponse
}

/* model for simplify = false
type TaobaoBusAgentMultipleRefundConfirmResponse struct {

    // 错误码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 失败错误信息
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 退款成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

}
*/

type TaobaoBusAgentMultipleRefundConfirmResponse struct {

    // 错误码
    ResultCode   string `json:"result_code,omitempty"`

    // 失败错误信息
    ResultMsg   string `json:"result_msg,omitempty"`

    // 退款成功
    IsSuccess   bool `json:"is_success,omitempty"`

}
