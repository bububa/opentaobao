package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池绑定接口 APIResponse
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口
*/
type AlibabaMarketingLotteryActivityBindAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_activity_bind_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关联成功
    
    Result   bool `json:"result,omitempty" xml:"