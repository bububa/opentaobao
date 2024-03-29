package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse 药品商家扫码 API返回值
// alibaba.alihealth.tracecodeplatform.code.entscan
//
// 药品商家扫描药品监管码，只有该商家的药才返回
type AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponseModel is 药品商家扫码 成功返回结果
type AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeplatform_code_entscan_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse
func GetAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse() *AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse {
	return poolAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse.Get().(*AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse 将 AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse(v *AlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodeplatformCodeEntscanAPIResponse.Put(v)
}
