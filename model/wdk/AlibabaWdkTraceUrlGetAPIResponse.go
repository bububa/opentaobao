package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkTraceUrlGetAPIResponse 溯源url透出 API返回值
// alibaba.wdk.trace.url.get
//
// 根据shopId和skuCode返回商品溯源url
type AlibabaWdkTraceUrlGetAPIResponse struct {
	model.CommonResponse
	AlibabaWdkTraceUrlGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkTraceUrlGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkTraceUrlGetAPIResponseModel).Reset()
}

// AlibabaWdkTraceUrlGetAPIResponseModel is 溯源url透出 成功返回结果
type AlibabaWdkTraceUrlGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_trace_url_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Data string `json:"data,omitempty" xml:"data,omitempty"`
	// code
	ReturnCode string `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkTraceUrlGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.Data = ""
	m.ReturnCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkTraceUrlGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkTraceUrlGetAPIResponse)
	},
}

// GetAlibabaWdkTraceUrlGetAPIResponse 从 sync.Pool 获取 AlibabaWdkTraceUrlGetAPIResponse
func GetAlibabaWdkTraceUrlGetAPIResponse() *AlibabaWdkTraceUrlGetAPIResponse {
	return poolAlibabaWdkTraceUrlGetAPIResponse.Get().(*AlibabaWdkTraceUrlGetAPIResponse)
}

// ReleaseAlibabaWdkTraceUrlGetAPIResponse 将 AlibabaWdkTraceUrlGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkTraceUrlGetAPIResponse(v *AlibabaWdkTraceUrlGetAPIResponse) {
	v.Reset()
	poolAlibabaWdkTraceUrlGetAPIResponse.Put(v)
}
