package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台查询可用奖品接口 API请求
alibaba.marketing.lottery.award.query

抽奖平台查询可用奖品接口
*/
type AlibabaMarketingLotteryAwardQueryRequest struct {
    model.Params
    // 查询奖品请求对象
    _lotteryAwardInstQuery   *LotteryAwardInstQueryDTO
}

// 初始化AlibabaMarketingLotteryAwardQueryRequest对象
func NewAlibabaMarketingLotteryAwardQueryRequest() *AlibabaMarketingLotteryAwardQueryRequest{
    return &AlibabaMarketingLotteryAwardQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryAwardQueryRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.award.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryAwardQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryAwardInstQuery Setter
// 查询奖品请求对象
func (r *AlibabaMarketingLotteryAwardQueryRequest) SetLotteryAwardInstQuery(_lotteryAwardInstQuery *LotteryAwardInstQueryDTO) error {
    r._lotteryAwardInstQuery = _lotteryAwardInstQuery
    r.Set("lottery_award_inst_query", _lotteryAwardInstQuery)
    return nil
}

// LotteryAwardInstQuery Getter
func (r AlibabaMarketingLotteryAwardQueryRequest) GetLotteryAwardInstQuery() *LotteryAwardInstQueryDTO {
    return r._lotteryAwardInstQuery
}
