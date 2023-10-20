package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabascbptargetadplanlistAPIResponse 定向推广-查询定向推广计划列表并返回计划基础信息 API返回值
// alibaba.scbp.target.ad.plan.list
//
// 定向推广-查询定向推广计划列表并返回计划基础信息
type AlibabascbptargetadplanlistAPIResponse struct {
	model.CommonResponse
	AlibabascbptargetadplanlistAPIResponseModel
}

// AlibabascbptargetadplanlistAPIResponseModel is 定向推广-查询定向推广计划列表并返回计划基础信息 成功返回结果
type AlibabascbptargetadplanlistAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 定向推广计划列表
	QuickCampaignList []TopP4pBasicQuickCampaignView `json:"quick_campaign_list,omitempty" xml:"quick_campaign_list>top_p4p_basic_quick_campaign_view,omitempty"`
	// 总页数
	TotalPage int64 `json:"total_page,omitempty" xml:"total_page,omitempty"`
	// 总个数
	TotalNum int64 `json:"total_num,omitempty" xml:"total_num,omitempty"`
}
