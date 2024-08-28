package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignAdd 信息流增加推广计划
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
func TaobaoFeedflowItemCampaignAdd(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignAddAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
