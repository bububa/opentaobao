package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse 查询无界版计划对应的原场景计划id API返回值
// taobao.universalbp.campaign.findsubcampaignid
//
// 查询该场景下，无界版计划对应的原场景的计划
type TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel).Reset()
}

// TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel is 查询无界版计划对应的原场景计划id 成功返回结果
type TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_campaign_findsubcampaignid_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpCampaignFindsubcampaignidTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpCampaignFindsubcampaignidAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse)
	},
}

// GetTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse
func GetTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse() *TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse {
	return poolTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse.Get().(*TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse)
}

// ReleaseTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse 将 TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse(v *TaobaoUniversalbpCampaignFindsubcampaignidAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpCampaignFindsubcampaignidAPIResponse.Put(v)
}
