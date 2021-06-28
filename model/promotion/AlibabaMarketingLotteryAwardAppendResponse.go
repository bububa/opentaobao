package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖品添加接口 APIResponse
alibaba.marketing.lottery.award.append

抽奖平台奖品添加接口，目前仅用于奖池众筹项目
*/
type AlibabaMarketingLotteryAwardAppendAPIResponse struct {
    model.CommonResponse
    AlibabaMarketingLotteryAwardAppendResponse
}

type AlibabaMarketingLotteryAwardAppendResponse struct {
    XMLName xml.Name `xml:"alibaba_marketing_lottery_award_append_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 奖品添加成功
    
    Result   bool `json:"result,omitempty" xml:"result,omitempty"`

    
    // code
    
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // 接口调用成功
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // msg
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
