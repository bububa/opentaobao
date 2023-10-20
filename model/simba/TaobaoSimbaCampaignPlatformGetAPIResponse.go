package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignPlatformGetAPIResponse 取得一个推广计划的投放平台设置 API返回值
// taobao.simba.campaign.platform.get
//
// 获得一个推广计划的投放平台设置
type TaobaoSimbaCampaignPlatformGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignPlatformGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignPlatformGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignPlatformGetAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignPlatformGetAPIResponseModel is 取得一个推广计划的投放平台设置 成功返回结果
type TaobaoSimbaCampaignPlatformGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_platform_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取得的推广计划的投放平台设置
	CampaignPlatform *CampaignPlatform `json:"campaign_platform,omitempty" xml:"campaign_platform,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignPlatformGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignPlatform = nil
}

var poolTaobaoSimbaCampaignPlatformGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignPlatformGetAPIResponse)
	},
}

// GetTaobaoSimbaCampaignPlatformGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignPlatformGetAPIResponse
func GetTaobaoSimbaCampaignPlatformGetAPIResponse() *TaobaoSimbaCampaignPlatformGetAPIResponse {
	return poolTaobaoSimbaCampaignPlatformGetAPIResponse.Get().(*TaobaoSimbaCampaignPlatformGetAPIResponse)
}

// ReleaseTaobaoSimbaCampaignPlatformGetAPIResponse 将 TaobaoSimbaCampaignPlatformGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignPlatformGetAPIResponse(v *TaobaoSimbaCampaignPlatformGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignPlatformGetAPIResponse.Put(v)
}
