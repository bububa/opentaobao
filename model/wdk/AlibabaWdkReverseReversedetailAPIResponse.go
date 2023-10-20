package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseReversedetailAPIResponse 退款详情 API返回值
// alibaba.wdk.reverse.reversedetail
//
// 退款详情
type AlibabaWdkReverseReversedetailAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseReversedetailAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkReverseReversedetailAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkReverseReversedetailAPIResponseModel).Reset()
}

// AlibabaWdkReverseReversedetailAPIResponseModel is 退款详情 成功返回结果
type AlibabaWdkReverseReversedetailAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_reversedetail_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkReverseReversedetailAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkReverseReversedetailAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseReversedetailAPIResponse)
	},
}

// GetAlibabaWdkReverseReversedetailAPIResponse 从 sync.Pool 获取 AlibabaWdkReverseReversedetailAPIResponse
func GetAlibabaWdkReverseReversedetailAPIResponse() *AlibabaWdkReverseReversedetailAPIResponse {
	return poolAlibabaWdkReverseReversedetailAPIResponse.Get().(*AlibabaWdkReverseReversedetailAPIResponse)
}

// ReleaseAlibabaWdkReverseReversedetailAPIResponse 将 AlibabaWdkReverseReversedetailAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkReverseReversedetailAPIResponse(v *AlibabaWdkReverseReversedetailAPIResponse) {
	v.Reset()
	poolAlibabaWdkReverseReversedetailAPIResponse.Put(v)
}
