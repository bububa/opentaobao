package simba

import (
    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的日限额 APIResponse
taobao.simba.campaign.budget.update

更新一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetUpdateAPIResponse struct {
    model.CommonResponse
    // Response *TaobaoSimbaCampaignBudgetUpdateResponse `json:"simba_campaign_budget_update_response,omitempty"` 
    TaobaoSimbaCampaignBudgetUpdateResponse
}

/* model for simplify = false
type TaobaoSimbaCampaignBudgetUpdateResponse struct {

    // 修改后的推广计划日限额
    
    CampaignBudget  *struct {
        CampaignBudget  *CampaignBudget `json:"campaign_budget,omitempty"`
    } `json:"campaign_budget,omitempty"`
    

}
*/

type TaobaoSimbaCampaignBudgetUpdateResponse struct {

    // 修改后的推广计划日限额
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty"`

}
