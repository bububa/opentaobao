package user

import (
    "github.com/bububa/opentaobao/model"
)

/* 
发奖接口 APIResponse
alibaba.benefit.send

发奖接口
*/
type AlibabaBenefitSendAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaBenefitSendResponse `json:"alibaba_benefit_send_response,omitempty"` 
    AlibabaBenefitSendResponse
}

/* model for simplify = false
type AlibabaBenefitSendResponse struct {

    // 接口返回消息
    
    ResultMsg   string `json:"result_msg,omitempty"`
    

    // 接口返回代码
    
    ResultCode   string `json:"result_code,omitempty"`
    

    // 是否处理成功
    
    ResultSuccess   bool `json:"result_success,omitempty"`
    

    // 权益id
    
    RightId   int64 `json:"right_id,omitempty"`
    

    // 奖品名称
    
    PrizeName   string `json:"prize_name,omitempty"`
    

}
*/

type AlibabaBenefitSendResponse struct {

    // 接口返回消息
    ResultMsg   string `json:"result_msg,omitempty"`

    // 接口返回代码
    ResultCode   string `json:"result_code,omitempty"`

    // 是否处理成功
    ResultSuccess   bool `json:"result_success,omitempty"`

    // 权益id
    RightId   int64 `json:"right_id,omitempty"`

    // 奖品名称
    PrizeName   string `json:"prize_name,omitempty"`

}
