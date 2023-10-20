package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignPage 批量查询计划列表
// taobao.feedflow.item.campaign.page
//
// 批量查询计划列表
func TaobaoFeedflowItemCampaignPage(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignPageAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignPageAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
