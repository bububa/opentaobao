package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryActivityQueryAPIRequest
抽奖平台奖池查询接口 API请求
alibaba.marketing.lottery.activity.query

抽奖平台奖池查询接口 */
type AlibabaMarketingLotteryActivityQueryAPIRequest struct {
	model.Params
	// 查询抽奖活动请求对象
	_lotteryActivityQuery *LotteryActivityQueryDto
}

// New
