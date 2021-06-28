package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台查询可用奖品接口 APIResponse
alibaba.marketing.lottery.award.query

抽奖平台查询可用奖品接口
*/
type AlibabaMarketingLotteryAwardQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryAwardQueryResponse `json:"alibaba_marketing_lottery_award_query_response,omitempty"` 
    AlibabaMarketingLotteryAwardQueryResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryAwardQueryResponse struct {

    // 返回结果
    
    Result  *struct {
        LotteryAwardInstResultDto  *LotteryAwardInstResultDto `json:"lottery_award_inst_result_dto,omitempty"`
    } `json:"result,omitempty"`
    

    // code
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msg
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryAwardQueryResponse struct {

    // 返回结果
    Result   *LotteryAwardInstResultDto `json:"result,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
