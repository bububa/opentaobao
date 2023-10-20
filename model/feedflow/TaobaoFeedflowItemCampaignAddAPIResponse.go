package feedflow

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaofeedflowitemcampaignaddAPIResponse 信息流增加推广计划 API返回值
// taobao.feedflow.item.campaign.add
//
// 信息流增加推广计划
type TaobaofeedflowitemcampaignaddAPIResponse struct {
	model.CommonResponse
	TaobaofeedflowitemcampaignaddAPIResponseModel
}

// TaobaofeedflowitemcampaignaddAPIResponseModel is 信息流增加推广计划 成功返回结果
type TaobaofeedflowitemcampaignaddAPIResponseModel struct {
	XMLName xml.Name `xml:"feedflow_item_campaign_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 异步获取历史数据接口返回结果
	Result *TaobaofeedflowitemcampaignaddResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
