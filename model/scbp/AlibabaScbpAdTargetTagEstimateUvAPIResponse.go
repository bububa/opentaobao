package scbp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaScbpAdTargetTagEstimateUvAPIResponse 标签人群预估 API返回值
// alibaba.scbp.ad.target.tag.estimate.uv
//
// 标签人群预估
type AlibabaScbpAdTargetTagEstimateUvAPIResponse struct {
	model.CommonResponse
	AlibabaScbpAdTargetTagEstimateUvAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagEstimateUvAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaScbpAdTargetTagEstimateUvAPIResponseModel).Reset()
}

// AlibabaScbpAdTargetTagEstimateUvAPIResponseModel is 标签人群预估 成功返回结果
type AlibabaScbpAdTargetTagEstimateUvAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_scbp_ad_target_tag_estimate_uv_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回数据：key:optionValue, value: 人群id
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaScbpAdTargetTagEstimateUvAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaScbpAdTargetTagEstimateUvAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaScbpAdTargetTagEstimateUvAPIResponse)
	},
}

// GetAlibabaScbpAdTargetTagEstimateUvAPIResponse 从 sync.Pool 获取 AlibabaScbpAdTargetTagEstimateUvAPIResponse
func GetAlibabaScbpAdTargetTagEstimateUvAPIResponse() *AlibabaScbpAdTargetTagEstimateUvAPIResponse {
	return poolAlibabaScbpAdTargetTagEstimateUvAPIResponse.Get().(*AlibabaScbpAdTargetTagEstimateUvAPIResponse)
}

// ReleaseAlibabaScbpAdTargetTagEstimateUvAPIResponse 将 AlibabaScbpAdTargetTagEstimateUvAPIResponse 保存到 sync.Pool
func ReleaseAlibabaScbpAdTargetTagEstimateUvAPIResponse(v *AlibabaScbpAdTargetTagEstimateUvAPIResponse) {
	v.Reset()
	poolAlibabaScbpAdTargetTagEstimateUvAPIResponse.Put(v)
}
