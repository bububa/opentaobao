package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖规则保存接口 APIRequest
alibaba.marketing.lottery.rule.save

抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则
*/
type AlibabaMarketingLotteryRuleSaveRequest struct {
    model.Params

    // 抽奖规则保存请求对象
    lotteryRuleCreate   *LotteryRuleCreateDto 

}

func NewAlibabaMarketingLotteryRuleSaveRequest() *AlibabaMarketingLotteryRuleSaveRequest{
    return &AlibabaMarketingLotteryRuleSaveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaMarketingLotteryRuleSaveRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.rule.save"
}

func (r AlibabaMarketingLotteryRuleSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaMarketingLotteryRuleSaveRequest) SetLotteryRuleCreate(lotteryRuleCreate *LotteryRuleCreateDto) error {
    r.lotteryRuleCreate = lotteryRuleCreate
    r.Set("lottery_rule_create", lotteryRuleCreate)
    return nil
}

func (r AlibabaMarketingLotteryRuleSaveRequest) GetLotteryRuleCreate() *LotteryRuleCreateDto {
    return r.lotteryRuleCreate
}

