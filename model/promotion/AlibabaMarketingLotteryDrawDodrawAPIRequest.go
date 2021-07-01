package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryDrawDodrawAPIRequest
抽奖平台抽奖接口 API请求
alibaba.marketing.lottery.draw.dodraw

抽奖平台PC端抽奖接口 */
type AlibabaMarketingLotteryDrawDodrawAPIRequest struct {
	model.Params
	// 抽奖请求对象
	_lotteryDrawQuery *LotteryDrawQueryDto
}

// New
