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
    Response *TaobaoSimbaCampaignBudgetGetResponse `json:"taobao_simba_campaign_budget_get_response,omitempty"`
}

type TaobaoSimbaCampaignBudgetGetResponse struct {

    // 推广计划日限额
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty"`

}
