package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse 获取场景计划的报表数据 API返回值
// taobao.onebp.dkx.campaign.campaign.reportpage
//
// 获取场景计划的报表数据。场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe，爆发收割adStrategyBaoFa
type TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel
}

// TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel is 获取场景计划的报表数据 成功返回结果
type TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_reportpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCampaignCampaignReportpageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
