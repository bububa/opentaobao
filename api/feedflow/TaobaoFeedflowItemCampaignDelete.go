package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignDelete 删除计划
// taobao.feedflow.item.campaign.delete
//
// 删除计划
func TaobaoFeedflowItemCampaignDelete(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignDeleteAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignDeleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
