package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseLoadFeatchorderAPIResponse 取货详情 API返回值
// alibaba.wdk.reverse.load.featchorder
//
// 取货详情
type AlibabaWdkReverseLoadFeatchorderAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseLoadFeatchorderAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkReverseLoadFeatchorderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkReverseLoadFeatchorderAPIResponseModel).Reset()
}

// AlibabaWdkReverseLoadFeatchorderAPIResponseModel is 取货详情 成功返回结果
type AlibabaWdkReverseLoadFeatchorderAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_load_featchorder_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ReverseResult
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkReverseLoadFeatchorderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkReverseLoadFeatchorderAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseLoadFeatchorderAPIResponse)
	},
}

// GetAlibabaWdkReverseLoadFeatchorderAPIResponse 从 sync.Pool 获取 AlibabaWdkReverseLoadFeatchorderAPIResponse
func GetAlibabaWdkReverseLoadFeatchorderAPIResponse() *AlibabaWdkReverseLoadFeatchorderAPIResponse {
	return poolAlibabaWdkReverseLoadFeatchorderAPIResponse.Get().(*AlibabaWdkReverseLoadFeatchorderAPIResponse)
}

// ReleaseAlibabaWdkReverseLoadFeatchorderAPIResponse 将 AlibabaWdkReverseLoadFeatchorderAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkReverseLoadFeatchorderAPIResponse(v *AlibabaWdkReverseLoadFeatchorderAPIResponse) {
	v.Reset()
	poolAlibabaWdkReverseLoadFeatchorderAPIResponse.Put(v)
}
