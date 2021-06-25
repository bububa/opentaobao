package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一组推广计划 APIResponse
taobao.simba.campaigns.get

取得一个客户的推广计划；
*/
type TaobaoSimbaCampaignsGetAPIResponse struct {
    model.CommonResponse
    Response *TaobaoSimbaCampaignsGetResponse `json:"taobao_simba_campaigns_get_response,omitempty"`
}

type TaobaoSimbaCampaignsGetResponse struct {

    // 推广计划列表
    Campaigns   []Campaign `json:"campaigns,omitempty"`

}
