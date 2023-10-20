package scs

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoonebpdkxreportreportcampaignrealtimeAPIResponse 查询某计划实时列表 API返回值
// taobao.onebp.dkx.report.report.campaign.realtime
//
// 查询某计划实时列表
// 入参1示例：{&#34;biz_code&#34;:&#34;adStrategyDkx&#34;}
// 入参2示例：{&#34;log_date_list&#34;: [&#34;2021-09-09&#34;],     &#34;campaign_id_list&#34;: [2821811599]}
type TaobaoonebpdkxreportreportcampaignrealtimeAPIResponse struct {
	model.CommonResponse
	TaobaoonebpdkxreportreportcampaignrealtimeAPIResponseModel
}

// TaobaoonebpdkxreportreportcampaignrealtimeAPIResponseModel is 查询某计划实时列表 成功返回结果
type TaobaoonebpdkxreportreportcampaignrealtimeAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_realtime_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoonebpdkxreportreportcampaignrealtimeResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
