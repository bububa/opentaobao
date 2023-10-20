package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignRptdailylist 推广计划分日数据查询
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
func TaobaoFeedflowItemCampaignRptdailylist(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignRptdailylistAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
