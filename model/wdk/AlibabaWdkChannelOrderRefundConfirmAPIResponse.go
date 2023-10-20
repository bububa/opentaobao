package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderRefundConfirmAPIResponse 退款确认 API返回值
// alibaba.wdk.channel.order.refund.confirm
//
// 退款确认
type AlibabaWdkChannelOrderRefundConfirmAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderRefundConfirmAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderRefundConfirmAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkChannelOrderRefundConfirmAPIResponseModel).Reset()
}

// AlibabaWdkChannelOrderRefundConfirmAPIResponseModel is 退款确认 成功返回结果
type AlibabaWdkChannelOrderRefundConfirmAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_refund_confirm_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderRefundConfirmApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderRefundConfirmAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkChannelOrderRefundConfirmAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderRefundConfirmAPIResponse)
	},
}

// GetAlibabaWdkChannelOrderRefundConfirmAPIResponse 从 sync.Pool 获取 AlibabaWdkChannelOrderRefundConfirmAPIResponse
func GetAlibabaWdkChannelOrderRefundConfirmAPIResponse() *AlibabaWdkChannelOrderRefundConfirmAPIResponse {
	return poolAlibabaWdkChannelOrderRefundConfirmAPIResponse.Get().(*AlibabaWdkChannelOrderRefundConfirmAPIResponse)
}

// ReleaseAlibabaWdkChannelOrderRefundConfirmAPIResponse 将 AlibabaWdkChannelOrderRefundConfirmAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkChannelOrderRefundConfirmAPIResponse(v *AlibabaWdkChannelOrderRefundConfirmAPIResponse) {
	v.Reset()
	poolAlibabaWdkChannelOrderRefundConfirmAPIResponse.Put(v)
}
