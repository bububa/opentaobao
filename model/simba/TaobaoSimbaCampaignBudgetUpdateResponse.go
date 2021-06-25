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
    Response *TaobaoSimbaCampaignBudgetUpdateResponse `json:"taobao_simba_campaign_budget_update_response,omitempty"`
}

type TaobaoSimbaCampaignBudgetUpdateResponse struct {

    // 修改后的推广计划日限额
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty"`

}
