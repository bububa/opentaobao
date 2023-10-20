package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderCreateAPIResponse 创建订单 API返回值
// alibaba.wdk.channel.order.create
//
// 外部商家创建订单
type AlibabaWdkChannelOrderCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkChannelOrderCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkChannelOrderCreateAPIResponseModel).Reset()
}

// AlibabaWdkChannelOrderCreateAPIResponseModel is 创建订单 成功返回结果
type AlibabaWdkChannelOrderCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_channel_order_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	ApiResult *AlibabaWdkChannelOrderCreateApiResult `json:"api_result,omitempty" xml:"api_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkChannelOrderCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiResult = nil
}

var poolAlibabaWdkChannelOrderCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkChannelOrderCreateAPIResponse)
	},
}

// GetAlibabaWdkChannelOrderCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkChannelOrderCreateAPIResponse
func GetAlibabaWdkChannelOrderCreateAPIResponse() *AlibabaWdkChannelOrderCreateAPIResponse {
	return poolAlibabaWdkChannelOrderCreateAPIResponse.Get().(*AlibabaWdkChannelOrderCreateAPIResponse)
}

// ReleaseAlibabaWdkChannelOrderCreateAPIResponse 将 AlibabaWdkChannelOrderCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkChannelOrderCreateAPIResponse(v *AlibabaWdkChannelOrderCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkChannelOrderCreateAPIResponse.Put(v)
}
