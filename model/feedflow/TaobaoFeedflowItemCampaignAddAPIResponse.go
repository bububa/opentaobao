package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignAddAPIResponse 信息流增加推广计划 API返回值
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
type TaobaoFeedflowItemCampaignAddAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCampaignAddAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCampaignAddAPIResponseModel is 信息流增加推广计划 成功返回结果
type TaobaoFeedflowItemCampaignAddAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaoFeedflowItemCampaignAddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCampaignAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignAddAPIResponse)
	},
}

// GetTaobaoFeedflowItemCampaignAddAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCampaignAddAPIResponse
func GetTaobaoFeedflowItemCampaignAddAPIResponse() *TaobaoFeedflowItemCampaignAddAPIResponse {
	return poolTaobaoFeedflowItemCampaignAddAPIResponse.Get().(*TaobaoFeedflowItemCampaignAddAPIResponse)
}

// ReleaseTaobaoFeedflowItemCampaignAddAPIResponse 将 TaobaoFeedflowItemCampaignAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignAddAPIResponse(v *TaobaoFeedflowItemCampaignAddAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignAddAPIResponse.Put(v)
}
