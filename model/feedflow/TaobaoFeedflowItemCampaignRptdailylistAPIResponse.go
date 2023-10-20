package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignRptdailylistAPIResponse 推广计划分日数据查询 API返回值
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
type TaobaoFeedflowItemCampaignRptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignRptdailylistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignRptdailylistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCampaignRptdailylistAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCampaignRptdailylistAPIResponseModel is 推广计划分日数据查询 成功返回结果
type TaobaoFeedflowItemCampaignRptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaoFeedflowItemCampaignRptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignRptdailylistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCampaignRptdailylistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignRptdailylistAPIResponse)
	},
}

// GetTaobaoFeedflowItemCampaignRptdailylistAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCampaignRptdailylistAPIResponse
func GetTaobaoFeedflowItemCampaignRptdailylistAPIResponse() *TaobaoFeedflowItemCampaignRptdailylistAPIResponse {
	return poolTaobaoFeedflowItemCampaignRptdailylistAPIResponse.Get().(*TaobaoFeedflowItemCampaignRptdailylistAPIResponse)
}

// ReleaseTaobaoFeedflowItemCampaignRptdailylistAPIResponse 将 TaobaoFeedflowItemCampaignRptdailylistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignRptdailylistAPIResponse(v *TaobaoFeedflowItemCampaignRptdailylistAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignRptdailylistAPIResponse.Put(v)
}
