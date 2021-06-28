package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
获取推广计划实时报表数据 APIResponse
taobao.simba.rtrpt.campaign.get

获取推广计划实时报表数据
*/
type TaobaoSimbaRtrptCampaignGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRtrptCampaignGetResponse `json:"simba_rtrpt_campaign_get_response,omitempty"` 
    TaobaoSimbaRtrptCampaignGetResponse
}

/* model for simplify = false
type TaobaoSimbaRtrptCampaignGetResponse struct {

    // 111
    
    Resultss  struct {
        RtRptResultEntityDTO  []RtRptResultEntityDTO `json:"rt_rpt_result_entity_dto,omitempty"`
    } `json:"resultss,omitempty"`
    

}
*/

type TaobaoSimbaRtrptCampaignGetResponse struct {

    // 111
    Resultss   []RtRptResultEntityDTO `json:"resultss,omitempty"`

}
