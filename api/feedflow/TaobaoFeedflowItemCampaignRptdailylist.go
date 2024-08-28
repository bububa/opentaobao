package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignRptdailylist 推广计划分日数据查询
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
func TaobaoFeedflowItemCampaignRptdailylist(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignRptdailylistAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignRptdailylistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
