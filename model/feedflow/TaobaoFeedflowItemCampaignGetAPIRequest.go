package feedflow

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoFeedflowItemCampaignGetAPIRequest
通过计划id查询计划 API请求
taobao.feedflow.item.campaign.get

通过计划id查询计划 */
type TaobaoFeedflowItemCampaignGetAPIRequest struct {
	model.Params
	// 计划id
	_campaginId int64
}

// New
