package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinAxbVendorPushCallEventAPIResponse 呼叫事件推送 API返回值
// alibaba.aliqin.axb.vendor.push.call.event
//
// 呼叫事件推送
// 响铃时间、摘机事件
type AlibabaAliqinAxbVendorPushCallEventAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinAxbVendorPushCallEventAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorPushCallEventAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinAxbVendorPushCallEventAPIResponseModel).Reset()
}

// AlibabaAliqinAxbVendorPushCallEventAPIResponseModel is 呼叫事件推送 成功返回结果
type AlibabaAliqinAxbVendorPushCallEventAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_axb_vendor_push_call_event_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 呼叫事件响应
	Result *AlibabaAliqinAxbVendorPushCallEventResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinAxbVendorPushCallEventAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinAxbVendorPushCallEventAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinAxbVendorPushCallEventAPIResponse)
	},
}

// GetAlibabaAliqinAxbVendorPushCallEventAPIResponse 从 sync.Pool 获取 AlibabaAliqinAxbVendorPushCallEventAPIResponse
func GetAlibabaAliqinAxbVendorPushCallEventAPIResponse() *AlibabaAliqinAxbVendorPushCallEventAPIResponse {
	return poolAlibabaAliqinAxbVendorPushCallEventAPIResponse.Get().(*AlibabaAliqinAxbVendorPushCallEventAPIResponse)
}

// ReleaseAlibabaAliqinAxbVendorPushCallEventAPIResponse 将 AlibabaAliqinAxbVendorPushCallEventAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinAxbVendorPushCallEventAPIResponse(v *AlibabaAliqinAxbVendorPushCallEventAPIResponse) {
	v.Reset()
	poolAlibabaAliqinAxbVendorPushCallEventAPIResponse.Put(v)
}
