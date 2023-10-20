package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse 查询某计划离线列表 API返回值
// taobao.onebp.dkx.report.report.campaign.offline
//
// 查询某计划离线列表；
// 拓展流量查询：
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;launch_product_id_list&#34;:[101004013],&#34;start_time&#34;:&#34;2021-04-26&#34;,&#34;campaign_id_list&#34;:[134821085],&#34;end_time&#34;:&#34;2021-04-28&#34;,&#34;effect&#34;:15,}
// 非拓展流量查询：
// 入参2示例：{&#34;start_time&#34;:&#34;2021-09-08&#34;,&#34;campaign_id_list&#34;:[2821811599],&#34;end_time&#34;:&#34;2021-09-08&#34;,&#34;effect&#34;:15}
type TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel is 查询某计划离线列表 成功返回结果
type TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCampaignOfflineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse
func GetTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse() *TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse {
	return poolTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse.Get().(*TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse 将 TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse(v *TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportCampaignOfflineAPIResponse.Put(v)
}
