package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotteryAwardAppendAPIRequest
抽奖平台奖品添加接口 API请求
alibaba.marketing.lottery.award.append

抽奖平台奖品添加接口，目前仅用于奖池众筹项目 */
type AlibabaMarketingLotteryAwardAppendAPIRequest struct {
	model.Params
	// 奖品添加请求对象
	_lotteryAwardAppend *LotteryAwardAppendDto
}

// New
