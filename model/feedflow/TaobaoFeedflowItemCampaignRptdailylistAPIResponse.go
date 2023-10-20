package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignrptdailylistAPIResponse 推广计划分日数据查询 API返回值
// taobao.feedflow.item.campaign.rptdailylist
//
// 推广计划分日数据查询
type TaobaofeedflowitemcampaignrptdailylistAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcampaignrptdailylistAPIResponseModel
}

// TaobaofeedflowitemcampaignrptdailylistAPIResponseModel is 推广计划分日数据查询 成功返回结果
type TaobaofeedflowitemcampaignrptdailylistAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_rptdailylist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *TaobaofeedflowitemcampaignrptdailylistResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
