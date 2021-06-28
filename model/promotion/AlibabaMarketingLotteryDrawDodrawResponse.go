package promotion

import (
    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖接口 APIResponse
alibaba.marketing.lottery.draw.dodraw

抽奖平台PC端抽奖接口
*/
type AlibabaMarketingLotteryDrawDodrawAPIResponse struct {
    model.CommonResponse
    // Response *AlibabaMarketingLotteryDrawDodrawResponse `json:"alibaba_marketing_lottery_draw_dodraw_response,omitempty"` 
    AlibabaMarketingLotteryDrawDodrawResponse
}

/* model for simplify = false
type AlibabaMarketingLotteryDrawDodrawResponse struct {

    // 抽奖结果
    
    LotteryDrawResult  *struct {
        LotteryDrawResultDto  *LotteryDrawResultDto `json:"lottery_draw_result_dto,omitempty"`
    } `json:"lottery_draw_result,omitempty"`
    

    // code
    
    MsgCode   int64 `json:"msg_code,omitempty"`
    

    // success
    
    IsSuccess   bool `json:"is_success,omitempty"`
    

    // msg
    
    MsgInfo   string `json:"msg_info,omitempty"`
    

}
*/

type AlibabaMarketingLotteryDrawDodrawResponse struct {

    // 抽奖结果
    LotteryDrawResult   *LotteryDrawResultDto `json:"lottery_draw_result,omitempty"`

    // code
    MsgCode   int64 `json:"msg_code,omitempty"`

    // success
    IsSuccess   bool `json:"is_success,omitempty"`

    // msg
    MsgInfo   string `json:"msg_info,omitempty"`

}
