package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的投放地域 APIResponse
taobao.simba.campaign.area.update

更新一个推广计划的投放地域
*/
type TaobaoSimbaCampaignAreaUpdateAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignAreaUpdateResponse `json:"taobao_simba_campaign_area_update_response,omitempty"`
}

type TaobaoSimbaCampaignAreaUpdateResponse struct {

    // 修改后的推广计划投放地域
    CampaignArea   *CampaignArea `json:"campaign_area,omitempty"`

}
