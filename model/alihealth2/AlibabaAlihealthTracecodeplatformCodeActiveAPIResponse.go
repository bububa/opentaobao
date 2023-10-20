package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse 正大鸡蛋激活追溯码 API返回值
// alibaba.alihealth.tracecodeplatform.code.active
//
// 用于正大鸡蛋激活追溯码
type AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodeplatformCodeActiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodeplatformCodeActiveAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodeplatformCodeActiveAPIResponseModel is 正大鸡蛋激活追溯码 成功返回结果
type AlibabaAlihealthTracecodeplatformCodeActiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeplatform_code_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *AlibabaAlihealthTracecodeplatformCodeActiveResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodeplatformCodeActiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse
func GetAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse() *AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse {
	return poolAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse.Get().(*AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse 将 AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse(v *AlibabaAlihealthTracecodeplatformCodeActiveAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodeplatformCodeActiveAPIResponse.Put(v)
}
