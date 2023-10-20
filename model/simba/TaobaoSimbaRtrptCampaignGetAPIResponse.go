package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRtrptCampaignGetAPIResponse 获取推广计划实时报表数据 API返回值
// taobao.simba.rtrpt.campaign.get
//
// 获取推广计划实时报表数据
type TaobaoSimbaRtrptCampaignGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRtrptCampaignGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptCampaignGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRtrptCampaignGetAPIResponseModel).Reset()
}

// TaobaoSimbaRtrptCampaignGetAPIResponseModel is 获取推广计划实时报表数据 成功返回结果
type TaobaoSimbaRtrptCampaignGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rtrpt_campaign_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 111
	Resultss []RtRptResultEntityDto `json:"resultss,omitempty" xml:"resultss>rt_rpt_result_entity_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRtrptCampaignGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Resultss = m.Resultss[:0]
}

var poolTaobaoSimbaRtrptCampaignGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRtrptCampaignGetAPIResponse)
	},
}

// GetTaobaoSimbaRtrptCampaignGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRtrptCampaignGetAPIResponse
func GetTaobaoSimbaRtrptCampaignGetAPIResponse() *TaobaoSimbaRtrptCampaignGetAPIResponse {
	return poolTaobaoSimbaRtrptCampaignGetAPIResponse.Get().(*TaobaoSimbaRtrptCampaignGetAPIResponse)
}

// ReleaseTaobaoSimbaRtrptCampaignGetAPIResponse 将 TaobaoSimbaRtrptCampaignGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRtrptCampaignGetAPIResponse(v *TaobaoSimbaRtrptCampaignGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRtrptCampaignGetAPIResponse.Put(v)
}
