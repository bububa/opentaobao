package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryRuleSaveAPIRequest
抽奖平台抽奖规则保存接口 API请求
alibaba.marketing.lottery.rule.save

抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则 */
type AlibabaMarketingLotteryRuleSaveAPIRequest struct {
	model.Params
	// 抽奖规则保存请求对象
	_lotteryRuleCreate *LotteryRuleCreateDto
}

// New
