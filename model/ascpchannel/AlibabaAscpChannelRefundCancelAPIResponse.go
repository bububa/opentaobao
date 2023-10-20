package ascpchannel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAscpChannelRefundCancelAPIResponse 渠道退款单撤销 API返回值
// alibaba.ascp.channel.refund.cancel
//
// 售后申请的撤回接口
type AlibabaAscpChannelRefundCancelAPIResponse struct {
	model.CommonResponse
	AlibabaAscpChannelRefundCancelAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAscpChannelRefundCancelAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAscpChannelRefundCancelAPIResponseModel).Reset()
}

// AlibabaAscpChannelRefundCancelAPIResponseModel is 渠道退款单撤销 成功返回结果
type AlibabaAscpChannelRefundCancelAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_ascp_channel_refund_cancel_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值包装,result为返回具体消息内容
	Result *ResultWrapper `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAscpChannelRefundCancelAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAscpChannelRefundCancelAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAscpChannelRefundCancelAPIResponse)
	},
}

// GetAlibabaAscpChannelRefundCancelAPIResponse 从 sync.Pool 获取 AlibabaAscpChannelRefundCancelAPIResponse
func GetAlibabaAscpChannelRefundCancelAPIResponse() *AlibabaAscpChannelRefundCancelAPIResponse {
	return poolAlibabaAscpChannelRefundCancelAPIResponse.Get().(*AlibabaAscpChannelRefundCancelAPIResponse)
}

// ReleaseAlibabaAscpChannelRefundCancelAPIResponse 将 AlibabaAscpChannelRefundCancelAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAscpChannelRefundCancelAPIResponse(v *AlibabaAscpChannelRefundCancelAPIResponse) {
	v.Reset()
	poolAlibabaAscpChannelRefundCancelAPIResponse.Put(v)
}
