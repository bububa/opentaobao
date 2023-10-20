package feedflow

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/feedflow"
)

// TaobaoFeedflowItemCampaignDaybudget 获取当日投放日预算总额
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
func TaobaoFeedflowItemCampaignDaybudget(clt *core.SDKClient, req *feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIRequest, session string) (*feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIResponse, error) {
	var resp feedflow.TaobaoFeedflowItemCampaignDaybudgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
