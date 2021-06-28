package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的日限额 APIResponse
taobao.simba.campaign.budget.get

取得一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetGetAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCampaignBudgetGetResponse `json:"simba_campaign_budget_get_response,omitempty"` 
    TaobaoSimbaCampaignBudgetGetResponse
}

/* model for simplify = false
type TaobaoSimbaCampaignBudgetGetResponse struct {

    // 推广计划日限额
    
    CampaignBudget  *struct {
        CampaignBudget  *CampaignBudget `json:"campaign_budget,omitempty"`
    } `json:"campaign_budget,omitempty"`
    

}
*/

type TaobaoSimbaCampaignBudgetGetResponse struct {

    // 推广计划日限额
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty"`

}
