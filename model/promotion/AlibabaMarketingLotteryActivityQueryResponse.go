package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池查询接口 APIResponse
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口
*/
type AlibabaMarketingLotteryActivityQueryAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryActivityQueryResponse `json:"alibaba_marketing_lottery_activity_query_response,omitempty"` 
    AlibabaMarketingLotteryActivityQueryResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryActivityQueryResponse struct {

    // 分页结果
    
    PagingDto  *struct {
        PagingDto  *PagingDto `json:"paging_dto,omitempty"`
    } `json:"paging_dto,omitempty"`
    

    // code
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msg
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryActivityQueryResponse struct {

    // 分页结果
    PagingDto   *PagingDto `json:"paging_dto,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
