package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖品添加接口 APIResponse
alibaba.marketing.lottery.award.append

抽奖平台奖品添加接口，目前仅用于奖池众筹项目
*/
type AlibabaMarketingLotteryAwardAppendAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryAwardAppendResponse `json:"alibaba_marketing_lottery_award_append_response,omitempty"` 
    AlibabaMarketingLotteryAwardAppendResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryAwardAppendResponse struct {

    // 奖品添加成功
    
    Result   bool `json:"result,omitempty"`
    

    // code
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // 接口调用成功
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msg
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryAwardAppendResponse struct {

    // 奖品添加成功
    Result   bool `json:"result,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // 接口调用成功
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
