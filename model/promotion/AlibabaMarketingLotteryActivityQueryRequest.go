package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池查询接口 API请求
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口
*/
type AlibabaMarketingLotteryActivityQueryRequest struct {
    model.Params
    // 查询抽奖活动请求对象
    lotteryActivityQuery   *LotteryActivityQueryDto
}

// 初始化AlibabaMarketingLotteryActivityQueryRequest对象
func NewAlibabaMarketingLotteryActivityQueryRequest() *AlibabaMarketingLotteryActivityQueryRequest{
    return &AlibabaMarketingLotteryActivityQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityQueryRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryActivityQuery Setter
// 查询抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityQueryRequest) SetLotteryActivityQuery(lotteryActivityQuery *LotteryActivityQueryDto) error {
    r.lotteryActivityQuery = lotteryActivityQuery
    r.Set("lottery_activity_query", lotteryActivityQuery)
    return nil
}

// LotteryActivityQuery Getter
func (r AlibabaMarketingLotteryActivityQueryRequest) GetLotteryActivityQuery() *LotteryActivityQueryDto {
    return r.lotteryActivityQuery
}
