package scbp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaScbpAdTargetTagEstimateUvAPIResponse
标签人群预估 API返回值
alibaba.scbp.ad.target.tag.estimate.uv

标签人群预估 */
type AlibabaScbpAdTargetTagEstimateUvAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdTargetTagEstimateUvAPIResponseModel
}

// AlibabaScbpAdTargetTagEstimateUvAPIResponseModel is 标签人群预估 成功返回结果
type AlibabaScbpAdTargetTagEstimateUvAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_estimate_uv_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据：key:optionValue, value: 人群id
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}
