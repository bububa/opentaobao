package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaHourReportCampaignGetAPIResponse 计划维度小时报表获取 API返回值
// taobao.simba.hour.report.campaign.get
//
// 计划维度小时报表获取
type TaobaoSimbaHourReportCampaignGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaHourReportCampaignGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaHourReportCampaignGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaHourReportCampaignGetAPIResponseModel).Reset()
}

// TaobaoSimbaHourReportCampaignGetAPIResponseModel is 计划维度小时报表获取 成功返回结果
type TaobaoSimbaHourReportCampaignGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_hour_report_campaign_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 11
	Results *RtRptResultEntityDto `json:"results,omitempty" xml:"results,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaHourReportCampaignGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = nil
}

var poolTaobaoSimbaHourReportCampaignGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaHourReportCampaignGetAPIResponse)
	},
}

// GetTaobaoSimbaHourReportCampaignGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaHourReportCampaignGetAPIResponse
func GetTaobaoSimbaHourReportCampaignGetAPIResponse() *TaobaoSimbaHourReportCampaignGetAPIResponse {
	return poolTaobaoSimbaHourReportCampaignGetAPIResponse.Get().(*TaobaoSimbaHourReportCampaignGetAPIResponse)
}

// ReleaseTaobaoSimbaHourReportCampaignGetAPIResponse 将 TaobaoSimbaHourReportCampaignGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaHourReportCampaignGetAPIResponse(v *TaobaoSimbaHourReportCampaignGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaHourReportCampaignGetAPIResponse.Put(v)
}
