package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖方案创建接口 APIResponse
alibaba.marketing.lottery.schema.create

抽奖平台抽奖方案创建接口
*/
type AlibabaMarketingLotterySchemaCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_schema_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 抽奖方案对象
    
    LotterySchema   *LotterySchemaDto `json:"lottery_schema,omitempty" xml:"