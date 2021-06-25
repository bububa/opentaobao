package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得推广计划的可设置投放地域列表 APIResponse
taobao.simba.campaign.areaoptions.get

取得推广计划的可设置投放地域列表
*/
type TaobaoSimbaCampaignAreaoptionsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignAreaoptionsGetResponse `json:"taobao_simba_campaign_areaoptions_get_response,omitempty"`
}

type TaobaoSimbaCampaignAreaoptionsGetResponse struct {

    // 推广计划所有可设置的投放地域
    AreaOptions   []AreaOption `json:"area_options,omitempty"`

}
