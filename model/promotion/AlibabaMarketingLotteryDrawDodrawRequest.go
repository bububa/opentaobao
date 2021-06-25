package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖接口 APIRequest
alibaba.marketing.lottery.draw.dodraw

抽奖平台PC端抽奖接口
*/
type AlibabaMarketingLotteryDrawDodrawRequest struct {
    model.Params

    // 抽奖请求对象
    lotteryDrawQuery   *LotteryDrawQueryDto 

}

func NewAlibabaMarketingLotteryDrawDodrawRequest() *AlibabaMarketingLotteryDrawDodrawRequest{
    return &AlibabaMarketingLotteryDrawDodrawRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryDrawDodrawRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.draw.dodraw"
}

func (r AlibabaMarketingLotteryDrawDodrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryDrawDodrawRequest) SetLotteryDrawQuery(lotteryDrawQuery *LotteryDrawQueryDto) error {
    r.lotteryDrawQuery = lotteryDrawQuery
    r.Set("lottery_draw_query", lotteryDrawQuery)
    return nil
}

func (r AlibabaMarketingLotteryDrawDodrawRequest) GetLotteryDrawQuery() *LotteryDrawQueryDto {
    return r.lotteryDrawQuery
}

