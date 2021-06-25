package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台查询可用奖品接口 APIRequest
alibaba.marketing.lottery.award.query

抽奖平台查询可用奖品接口
*/
type AlibabaMarketingLotteryAwardQueryRequest struct {
    model.Params

    // 查询奖品请求对象
    lotteryAwardInstQuery   *LotteryAwardInstQueryDto 

}

func NewAlibabaMarketingLotteryAwardQueryRequest() *AlibabaMarketingLotteryAwardQueryRequest{
    return &AlibabaMarketingLotteryAwardQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryAwardQueryRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.award.query"
}

func (r AlibabaMarketingLotteryAwardQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryAwardQueryRequest) SetLotteryAwardInstQuery(lotteryAwardInstQuery *LotteryAwardInstQueryDto) error {
    r.lotteryAwardInstQuery = lotteryAwardInstQuery
    r.Set("lottery_award_inst_query", lotteryAwardInstQuery)
    return nil
}

func (r AlibabaMarketingLotteryAwardQueryRequest) GetLotteryAwardInstQuery() *LotteryAwardInstQueryDto {
    return r.lotteryAwardInstQuery
}

