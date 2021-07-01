package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

/* AlibabaMarketingLotteryActivityUnbind
抽奖平台奖池解绑接口
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口 */
func AlibabaMarketingLotteryActivityUnbind(clt *core.SDKClient, req *promotion.AlibabaMarketingLotteryActivityUnbindAPIRequest, session string) (*promotion.AlibabaMarketingLotteryActivityUnbindAPIResponse, error) {
	var resp promotion.AlibabaMarketingLotteryActivityUnbindAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
