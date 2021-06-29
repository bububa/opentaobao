package promotion

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
抽奖平台抽奖规则保存接口 API请求
alibaba.marketing.lottery.rule.save

抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则
*/
type AlibabaMarketingLotteryRuleSaveRequest struct {
    model.Params
    // 抽奖规则保存请求对象
    _lotteryRuleCreate   *LotteryRuleCreateDTO
}

// 初始化AlibabaMarketingLotteryRuleSaveRequest对象
func NewAlibabaMarketingLotteryRuleSaveRequest() *AlibabaMarketingLotteryRuleSaveRequest{
    return &AlibabaMarketingLotteryRuleSaveRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryRuleSaveRequest) GetApiMethodName() string {
    return "alibaba.marketing.lottery.rule.save"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryRuleSaveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// LotteryRuleCreate Setter
// 抽奖规则保存请求对象
func (r *AlibabaMarketingLotteryRuleSaveRequest) SetLotteryRuleCreate(_lotteryRuleCreate *LotteryRuleCreateDTO) error {
    r._lotteryRuleCreate = _lotteryRuleCreate
    r.Set("lottery_rule_create", _lotteryRuleCreate)
    return nil
}

// LotteryRuleCreate Getter
func (r AlibabaMarketingLotteryRuleSaveRequest) GetLotteryRuleCreate() *LotteryRuleCreateDTO {
    return r._lotteryRuleCreate
}
