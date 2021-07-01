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
type AlibabaInteractLotterydrawDodrawAPIRequest struct {
    model.Params
    // 抽奖请求对象
    _lotteryDrawQuery   *LotteryDrawQueryDTO
}

// 初始化AlibabaInteractLotterydrawDodrawAPIRequest对象
func NewAlibabaInteractLotterydrawDodrawRequest() *AlibabaInteractLotterydrawDodrawAPIRequest{
    return &AlibabaInteractLotterydrawDodrawAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaInteractLotterydrawDodrawAPIRequest) GetApiMethodName() string {
    return "alibaba.interact.lotterydraw.dodraw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaInteractLotterydrawDodrawAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryDrawQuery Setter
// 抽奖请求对象
func (r *AlibabaInteractLotterydrawDodrawAPIRequest) SetLotteryDrawQuery(_lotteryDrawQuery *LotteryDrawQueryDTO) error {
    r._lotteryDrawQuery = _lotteryDrawQuery
    r.Set("lottery_draw_query", _lotteryDrawQuery)
    return nil
}

// LotteryDrawQuery Getter
func (r AlibabaInteractLotterydrawDodrawAPIRequest) GetLotteryDrawQuery() *LotteryDrawQueryDTO {
    return r._lotteryDrawQuery
}
