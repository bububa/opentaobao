package scbp

import (
    "github.com/bububa/opentaobao/model"
)

/* 
标签人群预估 APIResponse
alibaba.scbp.ad.target.tag.estimate.uv

标签人群预估
*/
type AlibabaScbpAdTargetTagEstimateUvAPIResponse struct {
    model.CommonResponse
    Response *AlibabaScbpAdTargetTagEstimateUvResponse `json:"alibaba_scbp_ad_target_tag_estimate_uv_response,omitempty"`
}

type AlibabaScbpAdTargetTagEstimateUvResponse struct {

    // 返回数据：key:optionValue, value: 人群id
    Result   string `json:"result,omitempty"`

}
