package simba

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaosubwaycampaignofflinefindAPIResponse 查询某计划离线多日汇总列表 API返回值
// taobao.subway.campaign.offline.find
//
// 查询某计划离线列表
type TaobaosubwaycampaignofflinefindAPIResponse struct {
	model.CommonResponse
	TaobaosubwaycampaignofflinefindAPIResponseModel
}

// TaobaosubwaycampaignofflinefindAPIResponseModel is 查询某计划离线多日汇总列表 成功返回结果
type TaobaosubwaycampaignofflinefindAPIResponseModel struct {
	XMLName xml.Name `xml:"subway_campaign_offline_find_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result []ReportResultTopDto `json:"result,omitempty" xml:"result>report_result_top_dto,omitempty"`
	// 提示
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 总条数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
