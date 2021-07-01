package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignAddAPIRequest
信息流增加推广计划 API请求
taobao.feedflow.item.campaign.add

信息流增加推广计划 */
type TaobaoFeedflowItemCampaignAddAPIRequest struct {
	model.Params
	// 计划信息
	_campaign *CampaignDto
}

// New
