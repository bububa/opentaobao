package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderUsercancelAPIResponse 用户发起售中取消 API返回值
// alibaba.wdk.channel.order.usercancel
//
// 用户发起售中取消
type AlibabaWdkChannelOrderUsercancelAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderUsercancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderUsercancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkChannelOrderUsercancelAPIResponseModel).Reset()
}

// AlibabaWdkChannelOrderUsercancelAPIResponseModel is 用户发起售中取消 成功返回结果
type AlibabaWdkChannelOrderUsercancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_usercancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderUsercancelApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderUsercancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkChannelOrderUsercancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderUsercancelAPIResponse)
	},
}

// GetAlibabaWdkChannelOrderUsercancelAPIResponse 从 sync.Pool 获取 AlibabaWdkChannelOrderUsercancelAPIResponse
func GetAlibabaWdkChannelOrderUsercancelAPIResponse() *AlibabaWdkChannelOrderUsercancelAPIResponse {
	return poolAlibabaWdkChannelOrderUsercancelAPIResponse.Get().(*AlibabaWdkChannelOrderUsercancelAPIResponse)
}

// ReleaseAlibabaWdkChannelOrderUsercancelAPIResponse 将 AlibabaWdkChannelOrderUsercancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkChannelOrderUsercancelAPIResponse(v *AlibabaWdkChannelOrderUsercancelAPIResponse) {
	v.Reset()
	poolAlibabaWdkChannelOrderUsercancelAPIResponse.Put(v)
}
