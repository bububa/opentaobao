package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryAwardQueryAPIRequest
抽奖平台查询可用奖品接口 API请求
alibaba.marketing.lottery.award.query

抽奖平台查询可用奖品接口 */
type AlibabaMarketingLotteryAwardQueryAPIRequest struct {
	model.Params
	// 查询奖品请求对象
	_lotteryAwardInstQuery *LotteryAwardInstQueryDto
}

// New
