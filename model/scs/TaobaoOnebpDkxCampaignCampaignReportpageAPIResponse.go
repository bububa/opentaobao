package scs

import (
	"encoding/xml"
	"sync"

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

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel).Reset()
}

// TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel is 获取场景计划的报表数据 成功返回结果
type TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_campaign_campaign_reportpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxCampaignCampaignReportpageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxCampaignCampaignReportpageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse)
	},
}

// GetTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse
func GetTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse() *TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse {
	return poolTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse.Get().(*TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse)
}

// ReleaseTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse 将 TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse(v *TaobaoOnebpDkxCampaignCampaignReportpageAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxCampaignCampaignReportpageAPIResponse.Put(v)
}
