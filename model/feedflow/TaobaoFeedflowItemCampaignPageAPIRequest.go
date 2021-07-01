package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignPageAPIRequest
批量查询计划列表 API请求
taobao.feedflow.item.campaign.page

批量查询计划列表 */
type TaobaoFeedflowItemCampaignPageAPIRequest struct {
	model.Params
	// 入参
	_campaignQuery *CampaignQueryDto
}

// New
