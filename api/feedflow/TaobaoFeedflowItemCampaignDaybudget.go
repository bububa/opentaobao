package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignDaybudget 获取当日投放日预算总额
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
func TaobaoFeedflowItemCampaignDaybudget(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIRequest, resp *feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
