package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoFeedflowItemCampaignGetAPIResponse 通过计划id查询计划 API返回值
// taobao.feedflow.item.campaign.get
//
// 通过计划id查询计划
type TaobaoFeedflowItemCampaignGetAPIResponse struct {
	model.CommonResponse
	TaobaoFeedflowItemCampaignGetAPIResponseModel
}

// TaobaoFeedflowItemCampaignGetAPIResponseModel is 通过计划id查询计划 成功返回结果
type TaobaoFeedflowItemCampaignGetAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaoFeedflowItemCampaignGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
