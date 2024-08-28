package feedflow

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignDaybudget 获取当日投放日预算总额
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
func TaobaoFeedflowItemCampaignDaybudget(ctx context.Context, clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
