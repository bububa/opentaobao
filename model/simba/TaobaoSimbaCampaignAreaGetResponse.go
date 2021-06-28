package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的投放地域设置 APIResponse
taobao.simba.campaign.area.get

取得一个推广计划的投放地域设置
*/
type TaobaoSimbaCampaignAreaGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCampaignAreaGetResponse `json:"simba_campaign_area_get_response,omitempty"` 
    TaobaoSimbaCampaignAreaGetResponse
}

/* model for simplify = false
type TaobaoSimbaCampaignAreaGetResponse struct {

    // 推广计划的投放地域配置
    
    CampaignArea  *struct {
        CampaignArea  *CampaignArea `json:"campaign_area,omitempty"`
    } `json:"campaign_area,omitempty"`
    

}
*/

type TaobaoSimbaCampaignAreaGetResponse struct {

    // 推广计划的投放地域配置
    CampaignArea   *CampaignArea `json:"campaign_area,omitempty"`

}
