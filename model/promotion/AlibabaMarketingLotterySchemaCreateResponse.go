package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖方案创建接口 APIResponse
alibaba.marketing.lottery.schema.create

抽奖平台抽奖方案创建接口
*/
type AlibabaMarketingLotterySchemaCreateAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotterySchemaCreateResponse `json:"alibaba_marketing_lottery_schema_create_response,omitempty"` 
    AlibabaMarketingLotterySchemaCreateResponse
}

/* model for simplify = false
type AlibabaMarketingLotterySchemaCreateResponse struct {

    // 抽奖方案对象
    
    LotterySchema  *struct {
        LotterySchemaDto  *LotterySchemaDto `json:"lottery_schema_dto,omitempty"`
    } `json:"lottery_schema,omitempty"`
    

    // code
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msg
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotterySchemaCreateResponse struct {

    // 抽奖方案对象
    LotterySchema   *LotterySchemaDto `json:"lottery_schema,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
