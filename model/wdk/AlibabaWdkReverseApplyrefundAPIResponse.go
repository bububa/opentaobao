package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseApplyrefundAPIResponse 逆向申请接口 API返回值
// alibaba.wdk.reverse.applyrefund
//
// 逆向渲染
type AlibabaWdkReverseApplyrefundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseApplyrefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkReverseApplyrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkReverseApplyrefundAPIResponseModel).Reset()
}

// AlibabaWdkReverseApplyrefundAPIResponseModel is 逆向申请接口 成功返回结果
type AlibabaWdkReverseApplyrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_applyrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回result
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkReverseApplyrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkReverseApplyrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseApplyrefundAPIResponse)
	},
}

// GetAlibabaWdkReverseApplyrefundAPIResponse 从 sync.Pool 获取 AlibabaWdkReverseApplyrefundAPIResponse
func GetAlibabaWdkReverseApplyrefundAPIResponse() *AlibabaWdkReverseApplyrefundAPIResponse {
	return poolAlibabaWdkReverseApplyrefundAPIResponse.Get().(*AlibabaWdkReverseApplyrefundAPIResponse)
}

// ReleaseAlibabaWdkReverseApplyrefundAPIResponse 将 AlibabaWdkReverseApplyrefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkReverseApplyrefundAPIResponse(v *AlibabaWdkReverseApplyrefundAPIResponse) {
	v.Reset()
	poolAlibabaWdkReverseApplyrefundAPIResponse.Put(v)
}
