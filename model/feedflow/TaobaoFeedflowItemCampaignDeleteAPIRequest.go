package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignDeleteAPIRequest
删除计划 API请求
taobao.feedflow.item.campaign.delete

删除计划 */
type TaobaoFeedflowItemCampaignDeleteAPIRequest struct {
	model.Params
	// 计划id
	_campaignId int64
}

// New
