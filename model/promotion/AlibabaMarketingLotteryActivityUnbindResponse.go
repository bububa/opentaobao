package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池解绑接口 APIResponse
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口
*/
type AlibabaMarketingLotteryActivityUnbindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_activity_unbind_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 解绑成功
    
    Result   bool `json:"result,omitempty" xml:"