package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台活动删除接口 APIRequest
alibaba.marketing.lottery.activity.delete

抽奖平台活动删除接口
*/
type AlibabaMarketingLotteryActivityDeleteRequest struct {
    model.Params

    // 抽奖活动删除对象
    lotteryActivityDelete   *LotteryActivityDeleteDto 

}

func NewAlibabaMarketingLotteryActivityDeleteRequest() *AlibabaMarketingLotteryActivityDeleteRequest{
    return &AlibabaMarketingLotteryActivityDeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryActivityDeleteRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.delete"
}

func (r AlibabaMarketingLotteryActivityDeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryActivityDeleteRequest) SetLotteryActivityDelete(lotteryActivityDelete *LotteryActivityDeleteDto) error {
    r.lotteryActivityDelete = lotteryActivityDelete
    r.Set("lottery_activity_delete", lotteryActivityDelete)
    return nil
}

func (r AlibabaMarketingLotteryActivityDeleteRequest) GetLotteryActivityDelete() *LotteryActivityDeleteDto {
    return r.lotteryActivityDelete
}

