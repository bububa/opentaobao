package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignModifyAPIRequest
信息流修改计划 API请求
taobao.feedflow.item.campaign.modify

信息流修改计划 */
type TaobaoFeedflowItemCampaignModifyAPIRequest struct {
	model.Params
	// 修改参数
	_campaign *CampaignDto
}

// New
