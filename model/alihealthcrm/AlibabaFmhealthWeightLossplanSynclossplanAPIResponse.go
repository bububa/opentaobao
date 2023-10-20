package alihealthcrm

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaFmhealthWeightLossplanSynclossplanAPIResponse 减重计划--同步减重计划 API返回值
// alibaba.fmhealth.weight.lossplan.synclossplan
//
// 减重计划--三方同步用户初始化减重计划给我们
type AlibabaFmhealthWeightLossplanSynclossplanAPIResponse struct {
	model.CommonResponse
	AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaFmhealthWeightLossplanSynclossplanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel).Reset()
}

// AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel is 减重计划--同步减重计划 成功返回结果
type AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_fmhealth_weight_lossplan_synclossplan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaFmhealthWeightLossplanSynclossplanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaFmhealthWeightLossplanSynclossplanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaFmhealthWeightLossplanSynclossplanAPIResponse)
	},
}

// GetAlibabaFmhealthWeightLossplanSynclossplanAPIResponse 从 sync.Pool 获取 AlibabaFmhealthWeightLossplanSynclossplanAPIResponse
func GetAlibabaFmhealthWeightLossplanSynclossplanAPIResponse() *AlibabaFmhealthWeightLossplanSynclossplanAPIResponse {
	return poolAlibabaFmhealthWeightLossplanSynclossplanAPIResponse.Get().(*AlibabaFmhealthWeightLossplanSynclossplanAPIResponse)
}

// ReleaseAlibabaFmhealthWeightLossplanSynclossplanAPIResponse 将 AlibabaFmhealthWeightLossplanSynclossplanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaFmhealthWeightLossplanSynclossplanAPIResponse(v *AlibabaFmhealthWeightLossplanSynclossplanAPIResponse) {
	v.Reset()
	poolAlibabaFmhealthWeightLossplanSynclossplanAPIResponse.Put(v)
}
