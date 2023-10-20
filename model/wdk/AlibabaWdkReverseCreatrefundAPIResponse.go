package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkReverseCreatrefundAPIResponse 逆向提交 API返回值
// alibaba.wdk.reverse.creatrefund
//
// 逆向申请提交
type AlibabaWdkReverseCreatrefundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkReverseCreatrefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkReverseCreatrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkReverseCreatrefundAPIResponseModel).Reset()
}

// AlibabaWdkReverseCreatrefundAPIResponseModel is 逆向提交 成功返回结果
type AlibabaWdkReverseCreatrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_reverse_creatrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// ReverseResult
	Result *ReverseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkReverseCreatrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaWdkReverseCreatrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkReverseCreatrefundAPIResponse)
	},
}

// GetAlibabaWdkReverseCreatrefundAPIResponse 从 sync.Pool 获取 AlibabaWdkReverseCreatrefundAPIResponse
func GetAlibabaWdkReverseCreatrefundAPIResponse() *AlibabaWdkReverseCreatrefundAPIResponse {
	return poolAlibabaWdkReverseCreatrefundAPIResponse.Get().(*AlibabaWdkReverseCreatrefundAPIResponse)
}

// ReleaseAlibabaWdkReverseCreatrefundAPIResponse 将 AlibabaWdkReverseCreatrefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkReverseCreatrefundAPIResponse(v *AlibabaWdkReverseCreatrefundAPIResponse) {
	v.Reset()
	poolAlibabaWdkReverseCreatrefundAPIResponse.Put(v)
}
