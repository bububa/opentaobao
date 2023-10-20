package alsc

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlscOrderOrderUploadAPIResponse 订单回流 API返回值
// alibaba.alsc.order.order.upload
//
// 第三方订单回流
type AlibabaAlscOrderOrderUploadAPIResponse struct {
	model.CommonResponse
	AlibabaAlscOrderOrderUploadAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAlscOrderOrderUploadAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAlscOrderOrderUploadAPIResponseModel).Reset()
}

// AlibabaAlscOrderOrderUploadAPIResponseModel is 订单回流 成功返回结果
type AlibabaAlscOrderOrderUploadAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_alsc_order_order_upload_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回包装类
	Result *BaseResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAlscOrderOrderUploadAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAlscOrderOrderUploadAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAlscOrderOrderUploadAPIResponse)
	},
}

// GetAlibabaAlscOrderOrderUploadAPIResponse 从 sync.Pool 获取 AlibabaAlscOrderOrderUploadAPIResponse
func GetAlibabaAlscOrderOrderUploadAPIResponse() *AlibabaAlscOrderOrderUploadAPIResponse {
	return poolAlibabaAlscOrderOrderUploadAPIResponse.Get().(*AlibabaAlscOrderOrderUploadAPIResponse)
}

// ReleaseAlibabaAlscOrderOrderUploadAPIResponse 将 AlibabaAlscOrderOrderUploadAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAlscOrderOrderUploadAPIResponse(v *AlibabaAlscOrderOrderUploadAPIResponse) {
	v.Reset()
	poolAlibabaAlscOrderOrderUploadAPIResponse.Put(v)
}
