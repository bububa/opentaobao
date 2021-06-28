package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池解绑接口 APIResponse
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口
*/
type AlibabaMarketingLotteryActivityUnbindAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryActivityUnbindResponse `json:"alibaba_marketing_lottery_activity_unbind_response,omitempty"` 
    AlibabaMarketingLotteryActivityUnbindResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryActivityUnbindResponse struct {

    // 解绑成功
    
    Result   bool `json:"result,omitempty"`
    

    // 错误码
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // 调用成功与否
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryActivityUnbindResponse struct {

    // 解绑成功
    Result   bool `json:"result,omitempty"`

    // 错误码
    MsgCode   int64 `json:"msg_code,omitempty"`

    // 调用成功与否
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

}
