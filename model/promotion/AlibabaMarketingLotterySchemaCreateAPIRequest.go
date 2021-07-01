package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaMarketingLotterySchemaCreateAPIRequest
抽奖平台抽奖方案创建接口 API请求
alibaba.marketing.lottery.schema.create

抽奖平台抽奖方案创建接口 */
type AlibabaMarketingLotterySchemaCreateAPIRequest struct {
	model.Params
	// 创建抽奖方案请求对象
	_lotterySchemaCreate *LotterySchemaCreateDto
}

// New
