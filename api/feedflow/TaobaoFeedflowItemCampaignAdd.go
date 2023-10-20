package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignAdd 信息流增加推广计划
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
func TaobaoFeedflowItemCampaignAdd(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignAddAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignAddAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
