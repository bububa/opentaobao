package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryActivityDeleteAPIRequest
抽奖平台活动删除接口 API请求
alibaba.marketing.lottery.activity.delete

抽奖平台活动删除接口 */
type AlibabaMarketingLotteryActivityDeleteAPIRequest struct {
	model.Params
	// 抽奖活动删除对象
	_lotteryActivityDelete *LotteryActivityDeleteDto
}

// New
