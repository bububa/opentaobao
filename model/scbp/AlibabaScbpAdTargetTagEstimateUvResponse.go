package scbp

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
标签人群预估 APIResponse
alibaba.scbp.ad.target.tag.estimate.uv

标签人群预估
*/
type AlibabaScbpAdTargetTagEstimateUvAPIResponse struct {
    model.CommonResponse
    AlibabaScbpAdTargetTagEstimateUvResponse
}

type AlibabaScbpAdTargetTagEstimateUvResponse struct {
    XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_estimate_uv_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 返回数据：key:optionValue, value: 人群id
    
    Result   string `json:"result,omitempty" xml:"result,omitempty"`

    
}
