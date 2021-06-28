package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
定向推广-人群标签ID获取(店铺老客、优选人群) APIResponse
alibaba.scbp.target.ad.plan.crowd.id.get

定向推广-人群标签ID获取(店铺老客、优选人群)
*/
type AlibabaScbpTargetAdPlanCrowdIdGetAPIResponse struct {
    model.CommonResponse
    AlibabaScbpTargetAdPlanCrowdIdGetResponse
}

type AlibabaScbpTargetAdPlanCrowdIdGetResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_target_ad_plan_crowd_id_get_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果list
    
    ResultList   []CrowdView `json:"result_list,omitempty" xml:"result_list>crowd_view,omitempty"`
    
    
}
