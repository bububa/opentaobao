package scs

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse 获取计划分日报表 API返回值
// taobao.onebp.dkx.report.report.campaign.daylist
//
// 获取计划分日报表，场景和bizCode的对应关系为：拉新快adStrategyDkx，上新快adStrategyShangXin ，货品加速adStrategyProductSpeed，入会快adStrategyRuHui，预热蓄水adStrategyYuRe
type TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse struct {
	model.CommonResponse
	TaobaoOnebpDkxReportReportCampaignDaylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOnebpDkxReportReportCampaignDaylistAPIResponseModel).Reset()
}

// TaobaoOnebpDkxReportReportCampaignDaylistAPIResponseModel is 获取计划分日报表 成功返回结果
type TaobaoOnebpDkxReportReportCampaignDaylistAPIResponseModel struct {
	XMLName xml.Name `xml:"onebp_dkx_report_report_campaign_daylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoOnebpDkxReportReportCampaignDaylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOnebpDkxReportReportCampaignDaylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse)
	},
}

// GetTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse 从 sync.Pool 获取 TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse
func GetTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse() *TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse {
	return poolTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse.Get().(*TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse)
}

// ReleaseTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse 将 TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse(v *TaobaoOnebpDkxReportReportCampaignDaylistAPIResponse) {
	v.Reset()
	poolTaobaoOnebpDkxReportReportCampaignDaylistAPIResponse.Put(v)
}
