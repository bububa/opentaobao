package simba

import (
	"encoding/xml"

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

// TaobaoSimbaCampaignTimetemplateFindAPIResponseModel is 获取分时折扣模板 成功返回结果
type TaobaoSimbaCampaignTimetemplateFindAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_timetemplate_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的推广组分页对象
	Templates []ADGroupPage `json:"templates,omitempty" xml:"templates>ad_group_page,omitempty"`
}
