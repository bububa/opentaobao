package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台活动删除接口 APIResponse
alibaba.marketing.lottery.activity.delete

抽奖平台活动删除接口
*/
type AlibabaMarketingLotteryActivityDeleteAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryActivityDeleteResponse `json:"alibaba_marketing_lottery_activity_delete_response,omitempty"` 
    AlibabaMarketingLotteryActivityDeleteResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryActivityDeleteResponse struct {

    // result
    
    Result   bool `json:"result,omitempty"`
    

    // code
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msg
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryActivityDeleteResponse struct {

    // result
    Result   bool `json:"result,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
