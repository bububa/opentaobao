package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractLotterydrawDodrawAPIRequest
无线端抽奖接口 API请求
alibaba.interact.lotterydraw.dodraw

商家抽奖平台无线端抽奖接口开放 */
type AlibabaInteractLotterydrawDodrawAPIRequest struct {
	model.Params
	// 抽奖请求对象
	_lotteryDrawQuery *LotteryDrawQueryDto
}

// New
