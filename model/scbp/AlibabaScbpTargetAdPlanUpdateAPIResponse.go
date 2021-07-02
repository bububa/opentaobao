package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanUpdateAPIResponse 定向推广-更新推广计划的基础信息 API返回值
// alibaba.scbp.target.ad.plan.update
//
// 定向推广-更新推广计划的基础信息
type AlibabaScbpTargetAdPlanUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanUpdateAPIResponseModel
}

// AlibabaScbpTargetAdPlanUpdateAPIResponseModel is 定向推广-更新推广计划的基础信息 成功返回结果
type AlibabaScbpTargetAdPlanUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// true修改成功，false修改失败
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}
