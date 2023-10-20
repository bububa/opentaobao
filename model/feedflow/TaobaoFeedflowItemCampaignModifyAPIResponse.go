package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignmodifyAPIResponse 信息流修改计划 API返回值
// taobao.feedflow.item.campaign.modify
//
// 信息流修改计划
type TaobaofeedflowitemcampaignmodifyAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcampaignmodifyAPIResponseModel
}

// TaobaofeedflowitemcampaignmodifyAPIResponseModel is 信息流修改计划 成功返回结果
type TaobaofeedflowitemcampaignmodifyAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *TaobaofeedflowitemcampaignmodifyResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
