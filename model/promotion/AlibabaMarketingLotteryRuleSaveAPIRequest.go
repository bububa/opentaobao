package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMarketingLotteryRuleSaveAPIRequest 抽奖平台抽奖规则保存接口 API请求
// alibaba.marketing.lottery.rule.save
//
// 抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则
type AlibabaMarketingLotteryRuleSaveAPIRequest struct {
	model.Params
	// 抽奖规则保存请求对象
	_lotteryRuleCreate *LotteryRuleCreateDto
}

// NewAlibabaMarketingLotteryRuleSaveRequest 初始化AlibabaMarketingLotteryRuleSaveAPIRequest对象
func NewAlibabaMarketingLotteryRuleSaveRequest() *AlibabaMarketingLotteryRuleSaveAPIRequest {
	return &AlibabaMarketingLotteryRuleSaveAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMarketingLotteryRuleSaveAPIRequest) GetApiMethodName() string {
	return "alibaba.marketing.lottery.rule.save"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMarketingLotteryRuleSaveAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetLotteryRuleCreate is LotteryRuleCreate Setter
// 抽奖规则保存请求对象
func (r *AlibabaMarketingLotteryRuleSaveAPIRequest) SetLotteryRuleCreate(_lotteryRuleCreate *LotteryRuleCreateDto) error {
	r._lotteryRuleCreate = _lotteryRuleCreate
	r.Set("lottery_rule_create", _lotteryRuleCreate)
	return nil
}

// GetLotteryRuleCreate LotteryRuleCreate Getter
func (r AlibabaMarketingLotteryRuleSaveAPIRequest) GetLotteryRuleCreate() *LotteryRuleCreateDto {
	return r._lotteryRuleCreate
}
