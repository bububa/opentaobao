package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse 销量明星跟新预算相关接口 API返回值
// taobao.simba.salestar.campaign.budget.update
//
// 更新一个推广计划的日限额
type TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponseModel is 销量明星跟新预算相关接口 成功返回结果
type TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_campaign_budget_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改后的推广计划日限额
	CampaignBudget *CampaignBudget `json:"campaign_budget,omitempty" xml:"campaign_budget,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CampaignBudget = nil
}

var poolTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse)
	},
}

// GetTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse
func GetTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse() *TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse {
	return poolTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse.Get().(*TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse)
}

// ReleaseTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse 将 TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse(v *TaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarCampaignBudgetUpdateAPIResponse.Put(v)
}
