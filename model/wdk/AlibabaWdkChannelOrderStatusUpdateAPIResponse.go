package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderStatusUpdateAPIResponse 订单状态变更 API返回值
// alibaba.wdk.channel.order.status.update
//
// 订单状态变更
type AlibabaWdkChannelOrderStatusUpdateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderStatusUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderStatusUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkChannelOrderStatusUpdateAPIResponseModel).Reset()
}

// AlibabaWdkChannelOrderStatusUpdateAPIResponseModel is 订单状态变更 成功返回结果
type AlibabaWdkChannelOrderStatusUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_status_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderStatusUpdateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderStatusUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkChannelOrderStatusUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderStatusUpdateAPIResponse)
	},
}

// GetAlibabaWdkChannelOrderStatusUpdateAPIResponse 从 sync.Pool 获取 AlibabaWdkChannelOrderStatusUpdateAPIResponse
func GetAlibabaWdkChannelOrderStatusUpdateAPIResponse() *AlibabaWdkChannelOrderStatusUpdateAPIResponse {
	return poolAlibabaWdkChannelOrderStatusUpdateAPIResponse.Get().(*AlibabaWdkChannelOrderStatusUpdateAPIResponse)
}

// ReleaseAlibabaWdkChannelOrderStatusUpdateAPIResponse 将 AlibabaWdkChannelOrderStatusUpdateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkChannelOrderStatusUpdateAPIResponse(v *AlibabaWdkChannelOrderStatusUpdateAPIResponse) {
	v.Reset()
	poolAlibabaWdkChannelOrderStatusUpdateAPIResponse.Put(v)
}
