package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderUserrefundAPIResponse 用户发起售后退款(整单/部分) API返回值
// alibaba.wdk.channel.order.userrefund
//
// 用户发起售后退款(整单/部分)
type AlibabaWdkChannelOrderUserrefundAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderUserrefundAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderUserrefundAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkChannelOrderUserrefundAPIResponseModel).Reset()
}

// AlibabaWdkChannelOrderUserrefundAPIResponseModel is 用户发起售后退款(整单/部分) 成功返回结果
type AlibabaWdkChannelOrderUserrefundAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_userrefund_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderUserrefundApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderUserrefundAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkChannelOrderUserrefundAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderUserrefundAPIResponse)
	},
}

// GetAlibabaWdkChannelOrderUserrefundAPIResponse 从 sync.Pool 获取 AlibabaWdkChannelOrderUserrefundAPIResponse
func GetAlibabaWdkChannelOrderUserrefundAPIResponse() *AlibabaWdkChannelOrderUserrefundAPIResponse {
	return poolAlibabaWdkChannelOrderUserrefundAPIResponse.Get().(*AlibabaWdkChannelOrderUserrefundAPIResponse)
}

// ReleaseAlibabaWdkChannelOrderUserrefundAPIResponse 将 AlibabaWdkChannelOrderUserrefundAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkChannelOrderUserrefundAPIResponse(v *AlibabaWdkChannelOrderUserrefundAPIResponse) {
	v.Reset()
	poolAlibabaWdkChannelOrderUserrefundAPIResponse.Put(v)
}
