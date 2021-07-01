package simba

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
取得一个推广计划的日限额 API返回值 
taobao.simba.campaign.budget.get

取得一个推广计划的日限额
*/
type TaobaoSimbaCampaignBudgetGetAPIResponse struct {
    model.CommonResponse
    TaobaoSimbaCampaignBudgetGetAPIResponseModel
}

// 取得一个推广计划的日限额 成功返回结果
type TaobaoSimbaCampaignBudgetGetAPIResponseModel struct {
    XMLName xml.Name `xml:"simba_campaign_budget_get_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 推广计划日限额
    CampaignBudget   *CampaignBudget `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`
}
