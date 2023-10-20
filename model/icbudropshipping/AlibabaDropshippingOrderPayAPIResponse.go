package icbudropshipping

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaDropshippingOrderPayAPIResponse alibaba dropshipping 支付代扣 API返回值
// alibaba.dropshipping.order.pay
//
// alibaba dropshipping 支付代扣
type AlibabaDropshippingOrderPayAPIResponse struct {
	model.CommonResponse
	AlibabaDropshippingOrderPayAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaDropshippingOrderPayAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaDropshippingOrderPayAPIResponseModel).Reset()
}

// AlibabaDropshippingOrderPayAPIResponseModel is alibaba dropshipping 支付代扣 成功返回结果
type AlibabaDropshippingOrderPayAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_dropshipping_order_pay_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response model
	Value *CashierPayResponse `json:"value,omitempty" xml:"value,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaDropshippingOrderPayAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Value = nil
}

var poolAlibabaDropshippingOrderPayAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaDropshippingOrderPayAPIResponse)
	},
}

// GetAlibabaDropshippingOrderPayAPIResponse 从 sync.Pool 获取 AlibabaDropshippingOrderPayAPIResponse
func GetAlibabaDropshippingOrderPayAPIResponse() *AlibabaDropshippingOrderPayAPIResponse {
	return poolAlibabaDropshippingOrderPayAPIResponse.Get().(*AlibabaDropshippingOrderPayAPIResponse)
}

// ReleaseAlibabaDropshippingOrderPayAPIResponse 将 AlibabaDropshippingOrderPayAPIResponse 保存到 sync.Pool
func ReleaseAlibabaDropshippingOrderPayAPIResponse(v *AlibabaDropshippingOrderPayAPIResponse) {
	v.Reset()
	poolAlibabaDropshippingOrderPayAPIResponse.Put(v)
}
