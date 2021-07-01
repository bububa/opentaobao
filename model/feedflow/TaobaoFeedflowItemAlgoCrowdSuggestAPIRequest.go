package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest
单品人群建议出价 API请求
taobao.feedflow.item.algo.crowd.suggest

给超级推荐的广告主查看建议出价 */
type TaobaoFeedflowItemAlgoCrowdSuggestAPIRequest struct {
	model.Params
	// 人群列表
	_crowds []CrowdDto
	// 预估的宝贝id
	_itemId int64
	// 预估的计划id
	_campaignId int64
}

// New
