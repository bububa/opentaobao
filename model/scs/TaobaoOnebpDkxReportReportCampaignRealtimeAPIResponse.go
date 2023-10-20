package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse 查询某计划实时列表 API返回值
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;log_date_list&#34;: [&#34;2021-09-09&#34;],     &#34;campaign_id_list&#34;: [2821811599]}
type TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel is 查询某计划实时列表 成功返回结果
type TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCampaignRealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse
func GetTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse() *TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse {
	return poolTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse.Get().(*TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse 将 TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse(v *TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse.Put(v)
}
