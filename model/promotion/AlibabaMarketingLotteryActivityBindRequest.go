package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池绑定接口 APIRequest
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口
*/
type AlibabaMarketingLotteryActivityBindRequest struct {
    model.Params

    // 关联抽奖活动请求对象
    lotteryActivityRel   *LotteryActivityRelDto 

}

func NewAlibabaMarketingLotteryActivityBindRequest() *AlibabaMarketingLotteryActivityBindRequest{
    return &AlibabaMarketingLotteryActivityBindRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryActivityBindRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.bind"
}

func (r AlibabaMarketingLotteryActivityBindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryActivityBindRequest) SetLotteryActivityRel(lotteryActivityRel *LotteryActivityRelDto) error {
    r.lotteryActivityRel = lotteryActivityRel
    r.Set("lottery_activity_rel", lotteryActivityRel)
    return nil
}

func (r AlibabaMarketingLotteryActivityBindRequest) GetLotteryActivityRel() *LotteryActivityRelDto {
    return r.lotteryActivityRel
}

