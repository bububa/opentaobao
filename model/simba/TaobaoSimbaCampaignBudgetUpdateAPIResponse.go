package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignBudgetUpdateAPIResponse 更新一个推广计划的日限额 API返回值
// taobao.simba.campaign.budget.update
//
// 更新一个推广计划的日限额
type TaobaoSimbaCampaignBudgetUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignBudgetUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignBudgetUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignBudgetUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignBudgetUpdateAPIResponseModel is 更新一个推广计划的日限额 成功返回结果
type TaobaoSimbaCampaignBudgetUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_budget_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划日限额
	CampaignBudget *CampaignBudget `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignBudgetUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignBudget = nil
}

var poolTaobaoSimbaCampaignBudgetUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignBudgetUpdateAPIResponse)
	},
}

// GetTaobaoSimbaCampaignBudgetUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignBudgetUpdateAPIResponse
func GetTaobaoSimbaCampaignBudgetUpdateAPIResponse() *TaobaoSimbaCampaignBudgetUpdateAPIResponse {
	return poolTaobaoSimbaCampaignBudgetUpdateAPIResponse.Get().(*TaobaoSimbaCampaignBudgetUpdateAPIResponse)
}

// ReleaseTaobaoSimbaCampaignBudgetUpdateAPIResponse 将 TaobaoSimbaCampaignBudgetUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignBudgetUpdateAPIResponse(v *TaobaoSimbaCampaignBudgetUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignBudgetUpdateAPIResponse.Put(v)
}
