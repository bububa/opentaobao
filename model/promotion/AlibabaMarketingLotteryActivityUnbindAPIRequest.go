package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryActivityUnbindAPIRequest
抽奖平台奖池解绑接口 API请求
alibaba.marketing.lottery.activity.unbind

抽奖平台奖池解绑接口 */
type AlibabaMarketingLotteryActivityUnbindAPIRequest struct {
	model.Params
	// 解绑抽奖活动请求对象
	_lotteryActivityRel *LotteryActivityRelDto
}

// New
