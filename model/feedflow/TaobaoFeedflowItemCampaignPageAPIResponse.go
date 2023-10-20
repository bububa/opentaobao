package feedflow

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignPageAPIResponse 批量查询计划列表 API返回值
// taobao.feedflow.item.campaign.page
//
// 批量查询计划列表
type TaobaoFeedflowItemCampaignPageAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignPageAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignPageAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoFeedflowItemCampaignPageAPIResponseModel).Reset()
}

// TaobaoFeedflowItemCampaignPageAPIResponseModel is 批量查询计划列表 成功返回结果
type TaobaoFeedflowItemCampaignPageAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_page_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemCampaignPageResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoFeedflowItemCampaignPageAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoFeedflowItemCampaignPageAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoFeedflowItemCampaignPageAPIResponse)
	},
}

// GetTaobaoFeedflowItemCampaignPageAPIResponse 从 sync.Pool 获取 TaobaoFeedflowItemCampaignPageAPIResponse
func GetTaobaoFeedflowItemCampaignPageAPIResponse() *TaobaoFeedflowItemCampaignPageAPIResponse {
	return poolTaobaoFeedflowItemCampaignPageAPIResponse.Get().(*TaobaoFeedflowItemCampaignPageAPIResponse)
}

// ReleaseTaobaoFeedflowItemCampaignPageAPIResponse 将 TaobaoFeedflowItemCampaignPageAPIResponse 保存到 sync.Pool
func ReleaseTaobaoFeedflowItemCampaignPageAPIResponse(v *TaobaoFeedflowItemCampaignPageAPIResponse) {
	v.Reset()
	poolTaobaoFeedflowItemCampaignPageAPIResponse.Put(v)
}
