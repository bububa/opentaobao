package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖接口 API请求
alibaba.marketing.lottery.draw.dodraw

抽奖平台PC端抽奖接口
*/
type AlibabaMarketingLotteryDrawDodrawRequest struct {
    model.Params
    // 抽奖请求对象
    _lotteryDrawQuery   *LotteryDrawQueryDTO
}

// 初始化AlibabaMarketingLotteryDrawDodrawRequest对象
func NewAlibabaMarketingLotteryDrawDodrawRequest() *AlibabaMarketingLotteryDrawDodrawRequest{
    return &AlibabaMarketingLotteryDrawDodrawRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryDrawDodrawRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.draw.dodraw"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryDrawDodrawRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryDrawQuery Setter
// 抽奖请求对象
func (r *AlibabaMarketingLotteryDrawDodrawRequest) SetLotteryDrawQuery(_lotteryDrawQuery *LotteryDrawQueryDTO) error {
    r._lotteryDrawQuery = _lotteryDrawQuery
    r.Set("lottery_draw_query", _lotteryDrawQuery)
    return nil
}

// LotteryDrawQuery Getter
func (r AlibabaMarketingLotteryDrawDodrawRequest) GetLotteryDrawQuery() *LotteryDrawQueryDTO {
    return r._lotteryDrawQuery
}
