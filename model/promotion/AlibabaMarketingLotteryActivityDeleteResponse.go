package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台活动删除接口 APIResponse
alibaba.marketing.lottery.activity.delete

抽奖平台活动删除接口
*/
type AlibabaMarketingLotteryActivityDeleteAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_activity_delete_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // result
    
    Result   bool `json:"result,omitempty" xml:"