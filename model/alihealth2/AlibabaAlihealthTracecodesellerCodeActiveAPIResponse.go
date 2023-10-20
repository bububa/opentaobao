package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerCodeActiveAPIResponse 上传激活码的文件 API返回值
// alibaba.alihealth.tracecodeseller.code.active
//
// 上传商品的激活码文件，存到系统中
type AlibabaAlihealthTracecodesellerCodeActiveAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerCodeActiveAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerCodeActiveAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerCodeActiveAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerCodeActiveAPIResponseModel is 上传激活码的文件 成功返回结果
type AlibabaAlihealthTracecodesellerCodeActiveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_code_active_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result string `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerCodeActiveAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = ""
}

var poolAlibabaAlihealthTracecodesellerCodeActiveAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerCodeActiveAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerCodeActiveAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerCodeActiveAPIResponse
func GetAlibabaAlihealthTracecodesellerCodeActiveAPIResponse() *AlibabaAlihealthTracecodesellerCodeActiveAPIResponse {
	return poolAlibabaAlihealthTracecodesellerCodeActiveAPIResponse.Get().(*AlibabaAlihealthTracecodesellerCodeActiveAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerCodeActiveAPIResponse 将 AlibabaAlihealthTracecodesellerCodeActiveAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerCodeActiveAPIResponse(v *AlibabaAlihealthTracecodesellerCodeActiveAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerCodeActiveAPIResponse.Put(v)
}
