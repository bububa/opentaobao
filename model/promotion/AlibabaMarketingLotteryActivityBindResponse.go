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
    AlibabaMarketingLotteryActivityBindResponse
}

type AlibabaMarketingLotteryActivityBindResponse struct {
    XMLName xml.Name `xml:"alibaba_marketing_lottery_activity_bind_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 关联成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // 错误码
    
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 是否调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // 错误信息
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
