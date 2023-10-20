package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignDaybudgetAPIResponse 获取当日投放日预算总额 API返回值
// taobao.feedflow.item.campaign.daybudget
//
// 获取当日投放日预算总额
type TaobaoFeedflowItemCampaignDaybudgetAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignDaybudgetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignDaybudgetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCampaignDaybudgetAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCampaignDaybudgetAPIResponseModel is 获取当日投放日预算总额 成功返回结果
type TaobaoFeedflowItemCampaignDaybudgetAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_daybudget_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemCampaignDaybudgetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignDaybudgetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCampaignDaybudgetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignDaybudgetAPIResponse)
	},
}

// GetTaobaoFeedflowItemCampaignDaybudgetAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCampaignDaybudgetAPIResponse
func GetTaobaoFeedflowItemCampaignDaybudgetAPIResponse() *TaobaoFeedflowItemCampaignDaybudgetAPIResponse {
	return poolTaobaoFeedflowItemCampaignDaybudgetAPIResponse.Get().(*TaobaoFeedflowItemCampaignDaybudgetAPIResponse)
}

// ReleaseTaobaoFeedflowItemCampaignDaybudgetAPIResponse 将 TaobaoFeedflowItemCampaignDaybudgetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignDaybudgetAPIResponse(v *TaobaoFeedflowItemCampaignDaybudgetAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignDaybudgetAPIResponse.Put(v)
}
