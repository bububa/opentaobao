package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignTimetemplateFindAPIResponse 获取分时折扣模板 API返回值
// taobao.simba.campaign.timetemplate.find
//
// 批量得到智能推广推广计划下的推广组
type TaobaoSimbaCampaignTimetemplateFindAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignTimetemplateFindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignTimetemplateFindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignTimetemplateFindAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignTimetemplateFindAPIResponseModel is 获取分时折扣模板 成功返回结果
type TaobaoSimbaCampaignTimetemplateFindAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_timetemplate_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的推广组分页对象
	Templates []ADGroupPage `json:"templates,omitempty" xml:"templates>ad_group_page,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignTimetemplateFindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Templates = m.Templates[:0]
}

var poolTaobaoSimbaCampaignTimetemplateFindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignTimetemplateFindAPIResponse)
	},
}

// GetTaobaoSimbaCampaignTimetemplateFindAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignTimetemplateFindAPIResponse
func GetTaobaoSimbaCampaignTimetemplateFindAPIResponse() *TaobaoSimbaCampaignTimetemplateFindAPIResponse {
	return poolTaobaoSimbaCampaignTimetemplateFindAPIResponse.Get().(*TaobaoSimbaCampaignTimetemplateFindAPIResponse)
}

// ReleaseTaobaoSimbaCampaignTimetemplateFindAPIResponse 将 TaobaoSimbaCampaignTimetemplateFindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignTimetemplateFindAPIResponse(v *TaobaoSimbaCampaignTimetemplateFindAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignTimetemplateFindAPIResponse.Put(v)
}
