package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的日限额 APIResponse
taobao.simba.campaign.budget.get

取得一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignBudgetGetResponse
}

type TaobaoSimbaCampaignBudgetGetResponse struct {
    XMLName xml.Name `xml:"simba_campaign_budget_get_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 推广计划日限额
    
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`

    
}
