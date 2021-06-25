package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线端抽奖接口 APIRequest
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放
*/
type AlibabaInteractLotterydrawDodrawRequest struct {
    model.Params

    // 抽奖请求对象
    lotteryDrawQuery   *LotteryDrawQueryDto 

}

func NewAlibabaInteractLotterydrawDodrawRequest() *AlibabaInteractLotterydrawDodrawRequest{
    return &AlibabaInteractLotterydrawDodrawRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaInteractLotterydrawDodrawRequest) GetApiMethodName() string {
    return "alibaba.interact.lotterydraw.dodraw"
}

func (r AlibabaInteractLotterydrawDodrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaInteractLotterydrawDodrawRequest) SetLotteryDrawQuery(lotteryDrawQuery *LotteryDrawQueryDto) error {
    r.lotteryDrawQuery = lotteryDrawQuery
    r.Set("lottery_draw_query", lotteryDrawQuery)
    return nil
}

func (r AlibabaInteractLotterydrawDodrawRequest) GetLotteryDrawQuery() *LotteryDrawQueryDto {
    return r.lotteryDrawQuery
}

