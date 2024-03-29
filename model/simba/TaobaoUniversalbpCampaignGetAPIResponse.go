package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignGetAPIResponse 查询单个计划详情 API返回值
// taobao.universalbp.campaign.get
//
// 查询单个计划详情信息（不包括报表数据）
type TaobaoUniversalbpCampaignGetAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCampaignGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaignGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCampaignGetAPIResponseModel).Reset()
}

// TaobaoUniversalbpCampaignGetAPIResponseModel is 查询单个计划详情 成功返回结果
type TaobaoUniversalbpCampaignGetAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaign_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCampaignGetTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaignGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCampaignGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignGetAPIResponse)
	},
}

// GetTaobaoUniversalbpCampaignGetAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCampaignGetAPIResponse
func GetTaobaoUniversalbpCampaignGetAPIResponse() *TaobaoUniversalbpCampaignGetAPIResponse {
	return poolTaobaoUniversalbpCampaignGetAPIResponse.Get().(*TaobaoUniversalbpCampaignGetAPIResponse)
}

// ReleaseTaobaoUniversalbpCampaignGetAPIResponse 将 TaobaoUniversalbpCampaignGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCampaignGetAPIResponse(v *TaobaoUniversalbpCampaignGetAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCampaignGetAPIResponse.Put(v)
}
