package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池创建接口 APIRequest
alibaba.marketing.lottery.activity.create

抽奖平台奖池创建接口
*/
type AlibabaMarketingLotteryActivityCreateRequest struct {
    model.Params

    // 抽奖活动创建请求对象
    lotteryActivityCreate   *LotteryActivityCreateDto 

}

func NewAlibabaMarketingLotteryActivityCreateRequest() *AlibabaMarketingLotteryActivityCreateRequest{
    return &AlibabaMarketingLotteryActivityCreateRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryActivityCreateRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.create"
}

func (r AlibabaMarketingLotteryActivityCreateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryActivityCreateRequest) SetLotteryActivityCreate(lotteryActivityCreate *LotteryActivityCreateDto) error {
    r.lotteryActivityCreate = lotteryActivityCreate
    r.Set("lottery_activity_create", lotteryActivityCreate)
    return nil
}

func (r AlibabaMarketingLotteryActivityCreateRequest) GetLotteryActivityCreate() *LotteryActivityCreateDto {
    return r.lotteryActivityCreate
}

