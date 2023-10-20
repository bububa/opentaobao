package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelMainRefundCreateAPIResponse 淘外分销逆向创单（未发货整单退） API返回值
// alibaba.ascp.channel.main.refund.create
//
// 淘外分销解决方案--订单--逆向创单（未发货整单退）
type AlibabaAscpChannelMainRefundCreateAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelMainRefundCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelMainRefundCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelMainRefundCreateAPIResponseModel).Reset()
}

// AlibabaAscpChannelMainRefundCreateAPIResponseModel is 淘外分销逆向创单（未发货整单退） 成功返回结果
type AlibabaAscpChannelMainRefundCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_main_refund_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelMainRefundCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelMainRefundCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelMainRefundCreateAPIResponse)
	},
}

// GetAlibabaAscpChannelMainRefundCreateAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelMainRefundCreateAPIResponse
func GetAlibabaAscpChannelMainRefundCreateAPIResponse() *AlibabaAscpChannelMainRefundCreateAPIResponse {
	return poolAlibabaAscpChannelMainRefundCreateAPIResponse.Get().(*AlibabaAscpChannelMainRefundCreateAPIResponse)
}

// ReleaseAlibabaAscpChannelMainRefundCreateAPIResponse 将 AlibabaAscpChannelMainRefundCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelMainRefundCreateAPIResponse(v *AlibabaAscpChannelMainRefundCreateAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelMainRefundCreateAPIResponse.Put(v)
}
