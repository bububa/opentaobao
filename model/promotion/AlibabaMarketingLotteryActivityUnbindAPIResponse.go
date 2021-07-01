package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池解绑接口 API返回值 
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口
*/
type AlibabaMarketingLotteryActivityUnbindAPIResponse struct {
    model.CommonResponse
    AlibabaMarketingLotteryActivityUnbindAPIResponseModel
}

// 抽奖平台奖池解绑接口 成功返回结果
type AlibabaMarketingLotteryActivityUnbindAPIResponseModel struct {
    XMLName xml.Name `xml:"alibaba_marketing_lottery_activity_unbind_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 解绑成功
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`
    // 错误码
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`
    // 调用成功与否
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
    // 错误信息
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`
}
