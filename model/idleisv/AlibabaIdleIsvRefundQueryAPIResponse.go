package idleisv

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIdleIsvRefundQueryAPIResponse 闲鱼已验货交易订单退款信息查询 API返回值
// alibaba.idle.isv.refund.query
//
// 闲鱼服务商交易订单退款信息查询-单个退款查询
type AlibabaIdleIsvRefundQueryAPIResponse struct {
	model.CommonResponse
	AlibabaIdleIsvRefundQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaIdleIsvRefundQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaIdleIsvRefundQueryAPIResponseModel).Reset()
}

// AlibabaIdleIsvRefundQueryAPIResponseModel is 闲鱼已验货交易订单退款信息查询 成功返回结果
type AlibabaIdleIsvRefundQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_idle_isv_refund_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaIdleIsvRefundQueryResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaIdleIsvRefundQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaIdleIsvRefundQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaIdleIsvRefundQueryAPIResponse)
	},
}

// GetAlibabaIdleIsvRefundQueryAPIResponse 从 sync.Pool 获取 AlibabaIdleIsvRefundQueryAPIResponse
func GetAlibabaIdleIsvRefundQueryAPIResponse() *AlibabaIdleIsvRefundQueryAPIResponse {
	return poolAlibabaIdleIsvRefundQueryAPIResponse.Get().(*AlibabaIdleIsvRefundQueryAPIResponse)
}

// ReleaseAlibabaIdleIsvRefundQueryAPIResponse 将 AlibabaIdleIsvRefundQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaIdleIsvRefundQueryAPIResponse(v *AlibabaIdleIsvRefundQueryAPIResponse) {
	v.Reset()
	poolAlibabaIdleIsvRefundQueryAPIResponse.Put(v)
}
