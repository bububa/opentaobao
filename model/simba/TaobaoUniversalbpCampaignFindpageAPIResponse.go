package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignFindpageAPIResponse 查询计划分页列表 API返回值
// taobao.universalbp.campaign.findpage
//
// 分页查询场景内的计划列表
type TaobaoUniversalbpCampaignFindpageAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCampaignFindpageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaignFindpageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCampaignFindpageAPIResponseModel).Reset()
}

// TaobaoUniversalbpCampaignFindpageAPIResponseModel is 查询计划分页列表 成功返回结果
type TaobaoUniversalbpCampaignFindpageAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaign_findpage_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCampaignFindpageTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaignFindpageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCampaignFindpageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignFindpageAPIResponse)
	},
}

// GetTaobaoUniversalbpCampaignFindpageAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCampaignFindpageAPIResponse
func GetTaobaoUniversalbpCampaignFindpageAPIResponse() *TaobaoUniversalbpCampaignFindpageAPIResponse {
	return poolTaobaoUniversalbpCampaignFindpageAPIResponse.Get().(*TaobaoUniversalbpCampaignFindpageAPIResponse)
}

// ReleaseTaobaoUniversalbpCampaignFindpageAPIResponse 将 TaobaoUniversalbpCampaignFindpageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCampaignFindpageAPIResponse(v *TaobaoUniversalbpCampaignFindpageAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCampaignFindpageAPIResponse.Put(v)
}
