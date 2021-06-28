package promotion

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖接口 APIResponse
alibaba.marketing.lottery.draw.dodraw

抽奖平台PC端抽奖接口
*/
type AlibabaMarketingLotteryDrawDodrawAPIResponse struct {
    model.CommonResponse
    AlibabaMarketingLotteryDrawDodrawResponse
}

type AlibabaMarketingLotteryDrawDodrawResponse struct {
    XMLName xml.Name `xml:"alibaba_marketing_lottery_draw_dodraw_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 抽奖结果
    
    LotteryDrawResult   *LotteryDrawResultDto `json:"lottery_draw_result,omitempty" xml:"lottery_draw_result,omitempty"`

    
    // code
    
    MsgCode   int64 `json:"msg_code,omitempty" xml:"msg_code,omitempty"`

    
    // success
    
    IsSuccess   bool `json:"is_success,omitempty" xml:"is_success,omitempty"`

    
    // msg
    
    MsgInfo   string `json:"msg_info,omitempty" xml:"msg_info,omitempty"`

    
}
