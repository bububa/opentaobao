package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新一个推广计划的日限额 APIResponse
taobao.simba.campaign.budget.update

更新一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetUpdateAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignBudgetUpdateResponse
}

type TaobaoSimbaCampaignBudgetUpdateResponse struct {
    XMLName xml.Name `xml:"simba_campaign_budget_update_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改后的推广计划日限额
    
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`

    
}
