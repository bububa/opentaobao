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
    Response *TaobaoSimbaRtrptCampaignGetResponse `json:"taobao_simba_rtrpt_campaign_get_response,omitempty"`
}

type TaobaoSimbaRtrptCampaignGetResponse struct {

    // 111
    Resultss   []RtRptResultEntityDTO `json:"resultss,omitempty"`

}
