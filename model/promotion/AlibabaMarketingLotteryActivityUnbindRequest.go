package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台奖池解绑接口 API请求
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口
*/
type AlibabaMarketingLotteryActivityUnbindRequest struct {
    model.Params
    // 解绑抽奖活动请求对象
    lotteryActivityRel   *LotteryActivityRelDto
}

// 初始化AlibabaMarketingLotteryActivityUnbindRequest对象
func NewAlibabaMarketingLotteryActivityUnbindRequest() *AlibabaMarketingLotteryActivityUnbindRequest{
    return &AlibabaMarketingLotteryActivityUnbindRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryActivityUnbindRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.activity.unbind"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryActivityUnbindRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryActivityRel Setter
// 解绑抽奖活动请求对象
func (r *AlibabaMarketingLotteryActivityUnbindRequest) SetLotteryActivityRel(lotteryActivityRel *LotteryActivityRelDto) error {
    r.lotteryActivityRel = lotteryActivityRel
    r.Set("lottery_activity_rel", lotteryActivityRel)
    return nil
}

// LotteryActivityRel Getter
func (r AlibabaMarketingLotteryActivityUnbindRequest) GetLotteryActivityRel() *LotteryActivityRelDto {
    return r.lotteryActivityRel
}
