package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台查询可用奖品接口 APIResponse
alibaba.marketing.lottery.award.query

抽奖平台查询可用奖品接口
*/
type AlibabaMarketingLotteryAwardQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_award_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回结果
    
    Result   *LotteryAwardInstResultDto `json:"result,omitempty" xml:"