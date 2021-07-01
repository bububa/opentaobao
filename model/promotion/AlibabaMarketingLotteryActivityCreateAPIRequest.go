package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryActivityCreateAPIRequest
抽奖平台奖池创建接口 API请求
alibaba.marketing.lottery.activity.create

抽奖平台奖池创建接口 */
type AlibabaMarketingLotteryActivityCreateAPIRequest struct {
	model.Params
	// 抽奖活动创建请求对象
	_lotteryActivityCreate *LotteryActivityCreateDto
}

// New
