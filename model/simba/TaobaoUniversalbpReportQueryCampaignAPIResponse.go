package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpReportQueryCampaignAPIResponse 计划报表查询 API返回值
// taobao.universalbp.report.query.campaign
//
// 计划报表查询
type TaobaoUniversalbpReportQueryCampaignAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpReportQueryCampaignAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryCampaignAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpReportQueryCampaignAPIResponseModel).Reset()
}

// TaobaoUniversalbpReportQueryCampaignAPIResponseModel is 计划报表查询 成功返回结果
type TaobaoUniversalbpReportQueryCampaignAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_report_query_campaign_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpReportQueryCampaignTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpReportQueryCampaignAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpReportQueryCampaignAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpReportQueryCampaignAPIResponse)
	},
}

// GetTaobaoUniversalbpReportQueryCampaignAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpReportQueryCampaignAPIResponse
func GetTaobaoUniversalbpReportQueryCampaignAPIResponse() *TaobaoUniversalbpReportQueryCampaignAPIResponse {
	return poolTaobaoUniversalbpReportQueryCampaignAPIResponse.Get().(*TaobaoUniversalbpReportQueryCampaignAPIResponse)
}

// ReleaseTaobaoUniversalbpReportQueryCampaignAPIResponse 将 TaobaoUniversalbpReportQueryCampaignAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpReportQueryCampaignAPIResponse(v *TaobaoUniversalbpReportQueryCampaignAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpReportQueryCampaignAPIResponse.Put(v)
}
