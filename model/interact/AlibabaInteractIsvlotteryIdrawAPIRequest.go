package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractIsvlotteryIdrawAPIRequest
互动到店抽奖 API请求
alibaba.interact.isvlottery.idraw

互动到店抽奖 */
type AlibabaInteractIsvlotteryIdrawAPIRequest struct {
	model.Params
	// 互动实例ID
	_interactId string
	// 抽奖ID列表 用逗号隔开
	_awardIds string
	// 埋点参数
	_ua string
}

// New
