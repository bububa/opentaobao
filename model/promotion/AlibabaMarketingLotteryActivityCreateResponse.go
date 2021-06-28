package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池创建接口 APIResponse
alibaba.marketing.lottery.activity.create

抽奖平台奖池创建接口
*/
type AlibabaMarketingLotteryActivityCreateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_activity_create_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 错误码
    
    MsgCode   int64 `json:"msg_code,omitempty" xml:"