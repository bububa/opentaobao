package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
创建一个推广计划 APIResponse
taobao.simba.campaign.add

创建一个推广计划
*/
type TaobaoSimbaCampaignAddAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignAddResponse `json:"taobao_simba_campaign_add_response,omitempty"`
}

type TaobaoSimbaCampaignAddResponse struct {

    // 创建的推广计划
    Campaign   *Campaign `json:"campaign,omitempty"`

}
