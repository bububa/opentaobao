package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse 查询某计划离线列表 API返回值
// taobao.onebp.dkx.report.report.campaign.offline
//
// 查询某计划离线列表；
// 拓展流量查询：
// 入参1示例：{"biz_code":"adStrategyDkx"}
// 入参2示例：{"launch_product_id_list":[101004013],"start_time":"2021-04-26","campaign_id_list":[134821085],"end_time":"2021-04-28","effect":15,}
// 非拓展流量查询：
// 入参2示例：{"start_time":"2021-09-08","campaign_id_list":[2821811599],"end_time":"2021-09-08","effect":15}
type TaobaoOnebpDkxReportReportCampaignOfflineAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel
}

// TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel is 查询某计划离线列表 成功返回结果
type TaobaoOnebpDkxReportReportCampaignOfflineAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_offline_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCampaignOfflineResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
