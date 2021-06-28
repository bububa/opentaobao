package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-计划开启/暂停/删除 APIResponse
alibaba.scbp.target.ad.plan.operation

定向推广-计划开启/暂停/删除
*/
type AlibabaScbpTargetAdPlanOperationAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanOperationResponse
}

type AlibabaScbpTargetAdPlanOperationResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_operation_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 修改成功记录数
    
    Result   int64 `json:"result,omitempty" xml:"result,omitempty"`

    
}
