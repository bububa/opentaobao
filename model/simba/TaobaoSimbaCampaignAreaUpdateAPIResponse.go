package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaUpdateAPIResponse 更新一个推广计划的投放地域 API返回值
// taobao.simba.campaign.area.update
//
// 更新一个推广计划的投放地域
type TaobaoSimbaCampaignAreaUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignAreaUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAreaUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignAreaUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignAreaUpdateAPIResponseModel is 更新一个推广计划的投放地域 成功返回结果
type TaobaoSimbaCampaignAreaUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_area_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划投放地域
	CampaignArea *CampaignArea `json:"campaign_area,omitempty" xml:"campaign_area,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAreaUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignArea = nil
}

var poolTaobaoSimbaCampaignAreaUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignAreaUpdateAPIResponse)
	},
}

// GetTaobaoSimbaCampaignAreaUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignAreaUpdateAPIResponse
func GetTaobaoSimbaCampaignAreaUpdateAPIResponse() *TaobaoSimbaCampaignAreaUpdateAPIResponse {
	return poolTaobaoSimbaCampaignAreaUpdateAPIResponse.Get().(*TaobaoSimbaCampaignAreaUpdateAPIResponse)
}

// ReleaseTaobaoSimbaCampaignAreaUpdateAPIResponse 将 TaobaoSimbaCampaignAreaUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignAreaUpdateAPIResponse(v *TaobaoSimbaCampaignAreaUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignAreaUpdateAPIResponse.Put(v)
}
