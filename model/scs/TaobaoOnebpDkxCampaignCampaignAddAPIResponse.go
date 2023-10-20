package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignAddAPIResponse 添加计划 API返回值
// taobao.onebp.dkx.campaign.campaign.add
//
// 添加计划
// 套餐包计划： {&#34;campaign_view&#34;:{&#34;ad_strategy_info&#34;:{&#34;budget&#34;:&#34;2000&#34;,&#34;item_ids&#34;:[617764134830],&#34;launch_strategy_type&#34;:&#34;cost_first&#34;,&#34;order_amount&#34;:&#34;2000&#34;},&#34;campaign_name&#34;:&#34;新建byTop测试&#34;,&#34;launch_time&#34;:{&#34;begin_time&#34;:&#34;2021-09-27 00:00:01&#34;,&#34;end_time&#34;:&#34;2021-09-31 00:00:01&#34;},&#34;marketing&#34;:{&#34;market_aim&#34;:1041,&#34;market_scene&#34;:1095},&#34;promotion_model&#34;:&#34;order&#34;}}
// 持续推广计划：
// {&#34;campaignViewDTO&#34;:{&#34;adStrategyInfo&#34;:{&#34;hide&#34;:false,&#34;itemIds&#34;:[45031793073],&#34;launchStrategyType&#34;:&#34;cost_first&#34;},&#34;campaignName&#34;:&#34;新建byTop&#34;,&#34;dayBudget&#34;:{&#34;dayBudget&#34;:&#34;100&#34;,&#34;dmcType&#34;:1},&#34;marketing&#34;:{&#34;marketAim&#34;:1037,&#34;marketScene&#34;:1023},&#34;promotionModel&#34;:&#34;daily&#34;}}
type TaobaoOnebpDkxCampaignCampaignAddAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel is 添加计划 成功返回结果
type TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCampaignCampaignAddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCampaignCampaignAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCampaignCampaignAddAPIResponse)
	},
}

// GetTaobaoOnebpDkxCampaignCampaignAddAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCampaignCampaignAddAPIResponse
func GetTaobaoOnebpDkxCampaignCampaignAddAPIResponse() *TaobaoOnebpDkxCampaignCampaignAddAPIResponse {
	return poolTaobaoOnebpDkxCampaignCampaignAddAPIResponse.Get().(*TaobaoOnebpDkxCampaignCampaignAddAPIResponse)
}

// ReleaseTaobaoOnebpDkxCampaignCampaignAddAPIResponse 将 TaobaoOnebpDkxCampaignCampaignAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCampaignCampaignAddAPIResponse(v *TaobaoOnebpDkxCampaignCampaignAddAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCampaignCampaignAddAPIResponse.Put(v)
}
