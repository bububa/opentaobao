package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* AlibabaMarketingLotteryRuleSave
抽奖平台抽奖规则保存接口
alibaba.marketing.lottery.rule.save

抽奖平台抽奖规则保存接口，对于同一主体，保存新规则会失效老的规则 */
func AlibabaMarketingLotteryRuleSave(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryRuleSaveAPIRequest, session string) (*promotion.AlibabaMarketingLotteryRuleSaveAPIResponse, error) {
	var resp promotion.AlibabaMarketingLotteryRuleSaveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
