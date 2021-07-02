package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpTargetAdPlanOperationAPIResponse 定向推广-计划开启/暂停/删除 API返回值
// alibaba.scbp.target.ad.plan.operation
//
// 定向推广-计划开启/暂停/删除
type AlibabaScbpTargetAdPlanOperationAPIResponse struct {
	model.CommonResponse
	AlibabaScbpTargetAdPlanOperationAPIResponseModel
}

// AlibabaScbpTargetAdPlanOperationAPIResponseModel is 定向推广-计划开启/暂停/删除 成功返回结果
type AlibabaScbpTargetAdPlanOperationAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_operation_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 修改成功记录数
	Result int64 `json:"result,omitempty" xml:"result,omitempty"`
}
