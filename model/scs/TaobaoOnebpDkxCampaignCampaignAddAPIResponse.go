package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignAddAPIResponse 添加计划 API返回值
// taobao.onebp.dkx.campaign.campaign.add
//
// 添加计划
// 套餐包计划： {"campaign_view":{"ad_strategy_info":{"budget":"2000","item_ids":[617764134830],"launch_strategy_type":"cost_first","order_amount":"2000"},"campaign_name":"新建byTop测试","launch_time":{"begin_time":"2021-09-27 00:00:01","end_time":"2021-09-31 00:00:01"},"marketing":{"market_aim":1041,"market_scene":1095},"promotion_model":"order"}}
// 持续推广计划：
// {"campaignViewDTO":{"adStrategyInfo":{"hide":false,"itemIds":[45031793073],"launchStrategyType":"cost_first"},"campaignName":"新建byTop","dayBudget":{"dayBudget":"100","dmcType":1},"marketing":{"marketAim":1037,"marketScene":1023},"promotionModel":"daily"}}
type TaobaoOnebpDkxCampaignCampaignAddAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel
}

// TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel is 添加计划 成功返回结果
type TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCampaignCampaignAddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
