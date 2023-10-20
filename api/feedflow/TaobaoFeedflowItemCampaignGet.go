package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignGet 通过计划id查询计划
// taobao.feedflow.item.campaign.get
//
// 通过计划id查询计划
func TaobaoFeedflowItemCampaignGet(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignGetAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
