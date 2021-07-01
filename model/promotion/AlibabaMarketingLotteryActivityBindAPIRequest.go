package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryActivityBindAPIRequest
抽奖平台奖池绑定接口 API请求
alibaba.marketing.lottery.activity.bind

抽奖平台奖池关联接口 */
type AlibabaMarketingLotteryActivityBindAPIRequest struct {
	model.Params
	// 关联抽奖活动请求对象
	_lotteryActivityRel *LotteryActivityRelDto
}

// New
