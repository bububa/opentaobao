package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池绑定接口 APIResponse
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口
*/
type AlibabaMarketingLotteryActivityBindAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryActivityBindResponse `json:"alibaba_marketing_lottery_activity_bind_response,omitempty"` 
    AlibabaMarketingLotteryActivityBindResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryActivityBindResponse struct {

    // 关联成功
    
    Result   bool `json:"result,omitempty"`
    

    // 错误码
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // 是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryActivityBindResponse struct {

    // 关联成功
    Result   bool `json:"result,omitempty"`

    // 错误码
    MsgCode   int64 `json:"msg_code,omitempty"`

    // 是否调用成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty"`

}
