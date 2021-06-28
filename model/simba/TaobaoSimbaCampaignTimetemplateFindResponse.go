package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
获取分时折扣模板 APIResponse
taobao.simba.campaign.timetemplate.find

批量得到智能推广推广计划下的推广组
*/
type TaobaoSimbaCampaignTimetemplateFindAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignTimetemplateFindResponse
}

type TaobaoSimbaCampaignTimetemplateFindResponse struct {
    XMLName xml.Name `xml:"simba_campaign_timetemplate_find_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回的推广组分页对象
    
    Templates   []ADGroupPage `json:"templates,omitempty" xml:"templates>ad_group_page,omitempty"`
    
    
}
