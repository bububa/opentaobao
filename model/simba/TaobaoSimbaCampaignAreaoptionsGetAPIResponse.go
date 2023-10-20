package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignAreaoptionsGetAPIResponse 取得推广计划的可设置投放地域列表 API返回值
// taobao.simba.campaign.areaoptions.get
//
// 取得推广计划的可设置投放地域列表
type TaobaoSimbaCampaignAreaoptionsGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAreaoptionsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel is 取得推广计划的可设置投放地域列表 成功返回结果
type TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_areaoptions_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划所有可设置的投放地域
	AreaOptions []AreaOption `json:"area_options,omitempty" xml:"area_options>area_option,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignAreaoptionsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.AreaOptions = m.AreaOptions[:0]
}

var poolTaobaoSimbaCampaignAreaoptionsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignAreaoptionsGetAPIResponse)
	},
}

// GetTaobaoSimbaCampaignAreaoptionsGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignAreaoptionsGetAPIResponse
func GetTaobaoSimbaCampaignAreaoptionsGetAPIResponse() *TaobaoSimbaCampaignAreaoptionsGetAPIResponse {
	return poolTaobaoSimbaCampaignAreaoptionsGetAPIResponse.Get().(*TaobaoSimbaCampaignAreaoptionsGetAPIResponse)
}

// ReleaseTaobaoSimbaCampaignAreaoptionsGetAPIResponse 将 TaobaoSimbaCampaignAreaoptionsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignAreaoptionsGetAPIResponse(v *TaobaoSimbaCampaignAreaoptionsGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignAreaoptionsGetAPIResponse.Put(v)
}
