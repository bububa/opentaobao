package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// TaobaoOnebpDkxCampaignCampaignAdd 添加计划
// taobao.onebp.dkx.campaign.campaign.add
//
// 添加计划
// 套餐包计划： {"campaign_view":{"ad_strategy_info":{"budget":"2000","item_ids":[617764134830],"launch_strategy_type":"cost_first","order_amount":"2000"},"campaign_name":"新建byTop测试","launch_time":{"begin_time":"2021-09-27 00:00:01","end_time":"2021-09-31 00:00:01"},"marketing":{"market_aim":1041,"market_scene":1095},"promotion_model":"order"}}
// 持续推广计划：
// {"campaignViewDTO":{"adStrategyInfo":{"hide":false,"itemIds":[45031793073],"launchStrategyType":"cost_first"},"campaignName":"新建byTop","dayBudget":{"dayBudget":"100","dmcType":1},"marketing":{"marketAim":1037,"marketScene":1023},"promotionModel":"daily"}}
func TaobaoOnebpDkxCampaignCampaignAdd(clt *core.SDKClient, req *scs.TaobaoOnebpDkxCampaignCampaignAddAPIRequest, session string) (*scs.TaobaoOnebpDkxCampaignCampaignAddAPIResponse, error) {
	var resp scs.TaobaoOnebpDkxCampaignCampaignAddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
