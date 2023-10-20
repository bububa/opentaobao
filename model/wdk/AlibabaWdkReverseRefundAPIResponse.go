package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseRefundAPIResponse 退款打款 API返回值
// alibaba.wdk.reverse.refund
//
// 五道口退款打款开放能力接口
type AlibabaWdkReverseRefundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseRefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkReverseRefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkReverseRefundAPIResponseModel).Reset()
}

// AlibabaWdkReverseRefundAPIResponseModel is 退款打款 成功返回结果
type AlibabaWdkReverseRefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_refund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaWdkReverseRefundResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkReverseRefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkReverseRefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseRefundAPIResponse)
	},
}

// GetAlibabaWdkReverseRefundAPIResponse 从 sync.Pool 获取 AlibabaWdkReverseRefundAPIResponse
func GetAlibabaWdkReverseRefundAPIResponse() *AlibabaWdkReverseRefundAPIResponse {
	return poolAlibabaWdkReverseRefundAPIResponse.Get().(*AlibabaWdkReverseRefundAPIResponse)
}

// ReleaseAlibabaWdkReverseRefundAPIResponse 将 AlibabaWdkReverseRefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkReverseRefundAPIResponse(v *AlibabaWdkReverseRefundAPIResponse) {
	v.Reset()
	poolAlibabaWdkReverseRefundAPIResponse.Put(v)
}
