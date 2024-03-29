package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignModifyAPIResponse 信息流修改计划 API返回值
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
type TaobaoFeedflowItemCampaignModifyAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCampaignModifyAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCampaignModifyAPIResponseModel is 信息流修改计划 成功返回结果
type TaobaoFeedflowItemCampaignModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemCampaignModifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCampaignModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignModifyAPIResponse)
	},
}

// GetTaobaoFeedflowItemCampaignModifyAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCampaignModifyAPIResponse
func GetTaobaoFeedflowItemCampaignModifyAPIResponse() *TaobaoFeedflowItemCampaignModifyAPIResponse {
	return poolTaobaoFeedflowItemCampaignModifyAPIResponse.Get().(*TaobaoFeedflowItemCampaignModifyAPIResponse)
}

// ReleaseTaobaoFeedflowItemCampaignModifyAPIResponse 将 TaobaoFeedflowItemCampaignModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignModifyAPIResponse(v *TaobaoFeedflowItemCampaignModifyAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignModifyAPIResponse.Put(v)
}
