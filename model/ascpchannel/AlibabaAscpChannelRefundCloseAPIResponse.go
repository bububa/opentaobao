package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundCloseAPIResponse 渠道退款单关闭 API返回值
// alibaba.ascp.channel.refund.close
//
// 渠道退款单关闭
type AlibabaAscpChannelRefundCloseAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelRefundCloseAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelRefundCloseAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelRefundCloseAPIResponseModel).Reset()
}

// AlibabaAscpChannelRefundCloseAPIResponseModel is 渠道退款单关闭 成功返回结果
type AlibabaAscpChannelRefundCloseAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_refund_close_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelRefundCloseAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelRefundCloseAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelRefundCloseAPIResponse)
	},
}

// GetAlibabaAscpChannelRefundCloseAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelRefundCloseAPIResponse
func GetAlibabaAscpChannelRefundCloseAPIResponse() *AlibabaAscpChannelRefundCloseAPIResponse {
	return poolAlibabaAscpChannelRefundCloseAPIResponse.Get().(*AlibabaAscpChannelRefundCloseAPIResponse)
}

// ReleaseAlibabaAscpChannelRefundCloseAPIResponse 将 AlibabaAscpChannelRefundCloseAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelRefundCloseAPIResponse(v *AlibabaAscpChannelRefundCloseAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelRefundCloseAPIResponse.Put(v)
}
