package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOrderPayResultQueryAPIResponse alibaba查询订单支付结果 API返回值
// alibaba.order.pay.result.query
//
// alibaba查询订单支付结果
type AlibabaOrderPayResultQueryAPIResponse struct {
	model.CommonResponse
	AlibabaOrderPayResultQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOrderPayResultQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOrderPayResultQueryAPIResponseModel).Reset()
}

// AlibabaOrderPayResultQueryAPIResponseModel is alibaba查询订单支付结果 成功返回结果
type AlibabaOrderPayResultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_order_pay_result_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// pay response
	Value *CashierPayResponse `json:"value,omitempty" xml:"value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOrderPayResultQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = nil
}

var poolAlibabaOrderPayResultQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOrderPayResultQueryAPIResponse)
	},
}

// GetAlibabaOrderPayResultQueryAPIResponse 从 sync.Pool 获取 AlibabaOrderPayResultQueryAPIResponse
func GetAlibabaOrderPayResultQueryAPIResponse() *AlibabaOrderPayResultQueryAPIResponse {
	return poolAlibabaOrderPayResultQueryAPIResponse.Get().(*AlibabaOrderPayResultQueryAPIResponse)
}

// ReleaseAlibabaOrderPayResultQueryAPIResponse 将 AlibabaOrderPayResultQueryAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOrderPayResultQueryAPIResponse(v *AlibabaOrderPayResultQueryAPIResponse) {
	v.Reset()
	poolAlibabaOrderPayResultQueryAPIResponse.Put(v)
}
