package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaRptCampaigneffectGetAPIResponse 推广计划效果报表数据对象 API返回值
// taobao.simba.rpt.campaigneffect.get
//
// 推广计划效果报表数据对象
type TaobaoSimbaRptCampaigneffectGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaRptCampaigneffectGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaRptCampaigneffectGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaRptCampaigneffectGetAPIResponseModel).Reset()
}

// TaobaoSimbaRptCampaigneffectGetAPIResponseModel is 推广计划效果报表数据对象 成功返回结果
type TaobaoSimbaRptCampaigneffectGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_rpt_campaigneffect_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划效果报表数据对象
	RptCampaignEffectList string `json:"rpt_campaign_effect_list,omitempty" xml:"rpt_campaign_effect_list,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaRptCampaigneffectGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RptCampaignEffectList = ""
}

var poolTaobaoSimbaRptCampaigneffectGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaRptCampaigneffectGetAPIResponse)
	},
}

// GetTaobaoSimbaRptCampaigneffectGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaRptCampaigneffectGetAPIResponse
func GetTaobaoSimbaRptCampaigneffectGetAPIResponse() *TaobaoSimbaRptCampaigneffectGetAPIResponse {
	return poolTaobaoSimbaRptCampaigneffectGetAPIResponse.Get().(*TaobaoSimbaRptCampaigneffectGetAPIResponse)
}

// ReleaseTaobaoSimbaRptCampaigneffectGetAPIResponse 将 TaobaoSimbaRptCampaigneffectGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaRptCampaigneffectGetAPIResponse(v *TaobaoSimbaRptCampaigneffectGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaRptCampaigneffectGetAPIResponse.Put(v)
}
