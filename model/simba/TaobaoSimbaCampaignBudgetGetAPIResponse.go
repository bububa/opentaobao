package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaCampaignBudgetGetAPIResponse 取得一个推广计划的日限额 API返回值
// taobao.simba.campaign.budget.get
//
// 取得一个推广计划的日限额
type TaobaoSimbaCampaignBudgetGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaCampaignBudgetGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignBudgetGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaCampaignBudgetGetAPIResponseModel).Reset()
}

// TaobaoSimbaCampaignBudgetGetAPIResponseModel is 取得一个推广计划的日限额 成功返回结果
type TaobaoSimbaCampaignBudgetGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_campaign_budget_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广计划日限额
	CampaignBudget *CampaignBudget `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaCampaignBudgetGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignBudget = nil
}

var poolTaobaoSimbaCampaignBudgetGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaCampaignBudgetGetAPIResponse)
	},
}

// GetTaobaoSimbaCampaignBudgetGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaCampaignBudgetGetAPIResponse
func GetTaobaoSimbaCampaignBudgetGetAPIResponse() *TaobaoSimbaCampaignBudgetGetAPIResponse {
	return poolTaobaoSimbaCampaignBudgetGetAPIResponse.Get().(*TaobaoSimbaCampaignBudgetGetAPIResponse)
}

// ReleaseTaobaoSimbaCampaignBudgetGetAPIResponse 将 TaobaoSimbaCampaignBudgetGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaCampaignBudgetGetAPIResponse(v *TaobaoSimbaCampaignBudgetGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaCampaignBudgetGetAPIResponse.Put(v)
}
