package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池创建接口 API请求
alibaba.marketing.lottery.activity.create

抽奖平台奖池创建接口
*/
type AlibabaMarketingLotteryActivityCreateRequest struct {
    model.Params
    // 抽奖活动创建请求对象
    lotteryActivityCreate   *LotteryActivityCreateDto
}

// 初始化AlibabaMarketingLotteryActivityCreateRequest对象
func NewAlibabaMarketingLotteryActivityCreateRequest() *AlibabaMarketingLotteryActivityCreateRequest{
    return &AlibabaMarketingLotteryActivityCreateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.create"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryActivityCreate Setter
// 抽奖活动创建请求对象
func (r *AlibabaMarketingLotteryActivityCreateRequest) SetLotteryActivityCreate(lotteryActivityCreate *LotteryActivityCreateDto) error {
    r.lotteryActivityCreate = lotteryActivityCreate
    r.Set("lottery_activity_create", lotteryActivityCreate)
    return nil
}

// LotteryActivityCreate Getter
func (r AlibabaMarketingLotteryActivityCreateRequest) GetLotteryActivityCreate() *LotteryActivityCreateDto {
    return r.lotteryActivityCreate
}
