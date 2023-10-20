package scs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/scs"
)

// Taobaoonebpdkxcampaigncampaignadd 添加计划
// taobao.onebp.dkx.campaign.campaign.add
//
// 添加计划
// 套餐包计划： {&#34;campaign_view&#34;:{&#34;ad_strategy_info&#34;:{&#34;budget&#34;:&#34;2000&#34;,&#34;item_ids&#34;:[617764134830],&#34;launch_strategy_type&#34;:&#34;cost_first&#34;,&#34;order_amount&#34;:&#34;2000&#34;},&#34;campaign_name&#34;:&#34;新建byTop测试&#34;,&#34;launch_time&#34;:{&#34;begin_time&#34;:&#34;2021-09-27 00:00:01&#34;,&#34;end_time&#34;:&#34;2021-09-31 00:00:01&#34;},&#34;marketing&#34;:{&#34;market_aim&#34;:1041,&#34;market_scene&#34;:1095},&#34;promotion_model&#34;:&#34;order&#34;}}
// 持续推广计划：
// {&#34;campaignViewDTO&#34;:{&#34;adStrategyInfo&#34;:{&#34;hide&#34;:false,&#34;itemIds&#34;:[45031793073],&#34;launchStrategyType&#34;:&#34;cost_first&#34;},&#34;campaignName&#34;:&#34;新建byTop&#34;,&#34;dayBudget&#34;:{&#34;dayBudget&#34;:&#34;100&#34;,&#34;dmcType&#34;:1},&#34;marketing&#34;:{&#34;marketAim&#34;:1037,&#34;marketScene&#34;:1023},&#34;promotionModel&#34;:&#34;daily&#34;}}
func Taobaoonebpdkxcampaigncampaignadd(clt *core.SDKClient, req *scs.TaobaoonebpdkxcampaigncampaignaddAPIRequest, session string) (*scs.TaobaoonebpdkxcampaigncampaignaddAPIResponse, error) {
	var resp scs.TaobaoonebpdkxcampaigncampaignaddAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
