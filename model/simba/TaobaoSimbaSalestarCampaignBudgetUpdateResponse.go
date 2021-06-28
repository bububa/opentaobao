package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
销量明星跟新预算相关接口 APIResponse
taobao.simba.salestar.campaign.budget.update

更新一个推广计划的日限额
*/
type TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaSalestarCampaignBudgetUpdateResponse `json:"simba_salestar_campaign_budget_update_response,omitempty"` 
    TaobaoSimbaSalestarCampaignBudgetUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaSalestarCampaignBudgetUpdateResponse struct {

    // 修改后的推广计划日限额
    
    CampaignBudget  *struct {
        CampaignBudget  *CampaignBudget `json:"campaign_budget,omitempty"`
    } `json:"campaign_budget,omitempty"`
    

}
*/

type TaobaoSimbaSalestarCampaignBudgetUpdateResponse struct {

    // 修改后的推广计划日限额
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty"`

}
