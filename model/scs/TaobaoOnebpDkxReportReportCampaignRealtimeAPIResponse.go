package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse 查询某计划实时列表 API返回值
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"log_date_list": ["2021-09-09"],     "campaign_id_list": [2821811599]}
type TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel
}

// TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel is 查询某计划实时列表 成功返回结果
type TaobaoOnebpDkxReportReportCampaignRealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCampaignRealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
