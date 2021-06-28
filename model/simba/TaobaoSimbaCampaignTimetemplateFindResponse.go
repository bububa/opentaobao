package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取分时折扣模板 APIResponse
taobao.simba.campaign.timetemplate.find

批量得到智能推广推广计划下的推广组
*/
type TaobaoSimbaCampaignTimetemplateFindAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCampaignTimetemplateFindResponse `json:"simba_campaign_timetemplate_find_response,omitempty"` 
    TaobaoSimbaCampaignTimetemplateFindResponse
}

/* model for simplify = false
type TaobaoSimbaCampaignTimetemplateFindResponse struct {

    // 返回的推广组分页对象
    
    Templates  struct {
        ADGroupPage  []ADGroupPage `json:"ad_group_page,omitempty"`
    } `json:"templates,omitempty"`
    

}
*/

type TaobaoSimbaCampaignTimetemplateFindResponse struct {

    // 返回的推广组分页对象
    Templates   []ADGroupPage `json:"templates,omitempty"`

}
