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
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_marketing_lottery_draw_dodraw_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 抽奖结果
    
    LotteryDrawResult   *LotteryDrawResultDto `json:"lottery_draw_result,omitempty" xml:"