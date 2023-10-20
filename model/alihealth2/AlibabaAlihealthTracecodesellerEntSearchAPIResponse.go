package alihealth2

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthTracecodesellerEntSearchAPIResponse 查询商家信息 API返回值
// alibaba.alihealth.tracecodeseller.ent.search
//
// 查询商家信息
type AlibabaAlihealthTracecodesellerEntSearchAPIResponse struct {
	model.CommonResponse
	AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerEntSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel).Reset()
}

// AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel is 查询商家信息 成功返回结果
type AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alihealth_tracecodeseller_ent_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TopResultModel `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlihealthTracecodesellerEntSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlihealthTracecodesellerEntSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlihealthTracecodesellerEntSearchAPIResponse)
	},
}

// GetAlibabaAlihealthTracecodesellerEntSearchAPIResponse 从 sync.Pool 获取 AlibabaAlihealthTracecodesellerEntSearchAPIResponse
func GetAlibabaAlihealthTracecodesellerEntSearchAPIResponse() *AlibabaAlihealthTracecodesellerEntSearchAPIResponse {
	return poolAlibabaAlihealthTracecodesellerEntSearchAPIResponse.Get().(*AlibabaAlihealthTracecodesellerEntSearchAPIResponse)
}

// ReleaseAlibabaAlihealthTracecodesellerEntSearchAPIResponse 将 AlibabaAlihealthTracecodesellerEntSearchAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlihealthTracecodesellerEntSearchAPIResponse(v *AlibabaAlihealthTracecodesellerEntSearchAPIResponse) {
	v.Reset()
	poolAlibabaAlihealthTracecodesellerEntSearchAPIResponse.Put(v)
}
