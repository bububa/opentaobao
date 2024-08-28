package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignPage 批量查询计划列表
// taobao.feedflow.item.campaign.page
//
// 批量查询计划列表
func TaobaoFeedflowItemCampaignPage(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignPageAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignPageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
