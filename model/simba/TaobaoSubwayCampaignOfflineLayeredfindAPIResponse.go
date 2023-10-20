package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSubwayCampaignOfflineLayeredfindAPIResponse 查询计划离线列表30天转化周期 API返回值
// taobao.subway.campaign.offline.layeredfind
//
// 查询某计划离线列表
type TaobaoSubwayCampaignOfflineLayeredfindAPIResponse struct {
	model.CommonResponse
	TaobaoSubwayCampaignOfflineLayeredfindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSubwayCampaignOfflineLayeredfindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSubwayCampaignOfflineLayeredfindAPIResponseModel).Reset()
}

// TaobaoSubwayCampaignOfflineLayeredfindAPIResponseModel is 查询计划离线列表30天转化周期 成功返回结果
type TaobaoSubwayCampaignOfflineLayeredfindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_campaign_offline_layeredfind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result []ReportResultTopDto `json:"result,omitempty" xml:"result>report_result_top_dto,omitempty"`
	// 提示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSubwayCampaignOfflineLayeredfindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = m.Result[:0]
	m.Message = ""
	m.TotalCount = 0
}

var poolTaobaoSubwayCampaignOfflineLayeredfindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSubwayCampaignOfflineLayeredfindAPIResponse)
	},
}

// GetTaobaoSubwayCampaignOfflineLayeredfindAPIResponse 从 sync.Pool 获取 TaobaoSubwayCampaignOfflineLayeredfindAPIResponse
func GetTaobaoSubwayCampaignOfflineLayeredfindAPIResponse() *TaobaoSubwayCampaignOfflineLayeredfindAPIResponse {
	return poolTaobaoSubwayCampaignOfflineLayeredfindAPIResponse.Get().(*TaobaoSubwayCampaignOfflineLayeredfindAPIResponse)
}

// ReleaseTaobaoSubwayCampaignOfflineLayeredfindAPIResponse 将 TaobaoSubwayCampaignOfflineLayeredfindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSubwayCampaignOfflineLayeredfindAPIResponse(v *TaobaoSubwayCampaignOfflineLayeredfindAPIResponse) {
	v.Reset()
	poolTaobaoSubwayCampaignOfflineLayeredfindAPIResponse.Put(v)
}
