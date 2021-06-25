package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池查询接口 APIRequest
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口
*/
type AlibabaMarketingLotteryActivityQueryRequest struct {
    model.Params

    // 查询抽奖活动请求对象
    lotteryActivityQuery   *LotteryActivityQueryDto 

}

func NewAlibabaMarketingLotteryActivityQueryRequest() *AlibabaMarketingLotteryActivityQueryRequest{
    return &AlibabaMarketingLotteryActivityQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryActivityQueryRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.query"
}

func (r AlibabaMarketingLotteryActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryActivityQueryRequest) SetLotteryActivityQuery(lotteryActivityQuery *LotteryActivityQueryDto) error {
    r.lotteryActivityQuery = lotteryActivityQuery
    r.Set("lottery_activity_query", lotteryActivityQuery)
    return nil
}

func (r AlibabaMarketingLotteryActivityQueryRequest) GetLotteryActivityQuery() *LotteryActivityQueryDto {
    return r.lotteryActivityQuery
}

