package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池解绑接口 APIRequest
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口
*/
type AlibabaMarketingLotteryActivityUnbindRequest struct {
    model.Params

    // 解绑抽奖活动请求对象
    lotteryActivityRel   *LotteryActivityRelDto 

}

func NewAlibabaMarketingLotteryActivityUnbindRequest() *AlibabaMarketingLotteryActivityUnbindRequest{
    return &AlibabaMarketingLotteryActivityUnbindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryActivityUnbindRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.unbind"
}

func (r AlibabaMarketingLotteryActivityUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryActivityUnbindRequest) SetLotteryActivityRel(lotteryActivityRel *LotteryActivityRelDto) error {
    r.lotteryActivityRel = lotteryActivityRel
    r.Set("lottery_activity_rel", lotteryActivityRel)
    return nil
}

func (r AlibabaMarketingLotteryActivityUnbindRequest) GetLotteryActivityRel() *LotteryActivityRelDto {
    return r.lotteryActivityRel
}

