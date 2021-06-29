package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
无线端抽奖接口 API请求
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放
*/
type AlibabaInteractLotterydrawDodrawRequest struct {
    model.Params
    // 抽奖请求对象
    lotteryDrawQuery   *LotteryDrawQueryDto
}

// 初始化AlibabaInteractLotterydrawDodrawRequest对象
func NewAlibabaInteractLotterydrawDodrawRequest() *AlibabaInteractLotterydrawDodrawRequest{
    return &AlibabaInteractLotterydrawDodrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractLotterydrawDodrawRequest) GetApiMethodName() string {
    return "alibaba.interact.lotterydraw.dodraw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractLotterydrawDodrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryDrawQuery Setter
// 抽奖请求对象
func (r *AlibabaInteractLotterydrawDodrawRequest) SetLotteryDrawQuery(lotteryDrawQuery *LotteryDrawQueryDto) error {
    r.lotteryDrawQuery = lotteryDrawQuery
    r.Set("lottery_draw_query", lotteryDrawQuery)
    return nil
}

// LotteryDrawQuery Getter
func (r AlibabaInteractLotterydrawDodrawRequest) GetLotteryDrawQuery() *LotteryDrawQueryDto {
    return r.lotteryDrawQuery
}
