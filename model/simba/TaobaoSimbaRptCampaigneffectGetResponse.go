package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
推广计划效果报表数据对象 APIResponse
taobao.simba.rpt.campaigneffect.get

推广计划效果报表数据对象
*/
type TaobaoSimbaRptCampaigneffectGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaRptCampaigneffectGetResponse `json:"simba_rpt_campaigneffect_get_response,omitempty"` 
    TaobaoSimbaRptCampaigneffectGetResponse
}

/* model for simplify = false
type TaobaoSimbaRptCampaigneffectGetResponse struct {

    // 推广计划效果报表数据对象
    
    RptCampaignEffectList   string `json:"rpt_campaign_effect_list,omitempty"`
    

}
*/

type TaobaoSimbaRptCampaigneffectGetResponse struct {

    // 推广计划效果报表数据对象
    RptCampaignEffectList   string `json:"rpt_campaign_effect_list,omitempty"`

}
