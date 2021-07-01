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
type AlibabaMarketingLotteryActivityQueryAPIRequest struct {
    model.Params
    // 查询抽奖活动请求对象
    _lotteryActivityQuery   *LotteryActivityQueryDTO
}

// 初始化AlibabaMarketingLotteryActivityQueryAPIRequest对象
func NewAlibabaMarketingLotteryActivityQueryRequest() *AlibabaMarketingLotteryActivityQueryAPIRequest{
    return &AlibabaMarketingLotteryActivityQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryActivityQuery Setter
// 查询抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityQueryAPIRequest) SetLotteryActivityQuery(_lotteryActivityQuery *LotteryActivityQueryDTO) error {
    r._lotteryActivityQuery = _lotteryActivityQuery
    r.Set("lottery_activity_query", _lotteryActivityQuery)
    return nil
}

// LotteryActivityQuery Getter
func (r AlibabaMarketingLotteryActivityQueryAPIRequest) GetLotteryActivityQuery() *LotteryActivityQueryDTO {
    return r._lotteryActivityQuery
}
