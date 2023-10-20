package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaGetAPIResponse 取得一个推广计划的投放地域设置 API返回值
// taobao.simba.campaign.area.get
//
// 取得一个推广计划的投放地域设置
type TaobaoSimbaCampaignAreaGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignAreaGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAreaGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignAreaGetAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignAreaGetAPIResponseModel is 取得一个推广计划的投放地域设置 成功返回结果
type TaobaoSimbaCampaignAreaGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_area_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划的投放地域配置
	CampaignArea *CampaignArea `json:"campaign_area,omitempty" xml:"campaign_area,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAreaGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignArea = nil
}

var poolTaobaoSimbaCampaignAreaGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignAreaGetAPIResponse)
	},
}

// GetTaobaoSimbaCampaignAreaGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignAreaGetAPIResponse
func GetTaobaoSimbaCampaignAreaGetAPIResponse() *TaobaoSimbaCampaignAreaGetAPIResponse {
	return poolTaobaoSimbaCampaignAreaGetAPIResponse.Get().(*TaobaoSimbaCampaignAreaGetAPIResponse)
}

// ReleaseTaobaoSimbaCampaignAreaGetAPIResponse 将 TaobaoSimbaCampaignAreaGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignAreaGetAPIResponse(v *TaobaoSimbaCampaignAreaGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignAreaGetAPIResponse.Put(v)
}
