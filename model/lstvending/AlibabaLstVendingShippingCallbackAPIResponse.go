package lstvending

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaLstVendingShippingCallbackAPIResponse 售货机出货回传接口 API返回值
// alibaba.lst.vending.shipping.callback
//
// 零售通自动售货机商品出货回传接口，同步商品出库最新状态。
type AlibabaLstVendingShippingCallbackAPIResponse struct {
	model.CommonResponse
	AlibabaLstVendingShippingCallbackAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaLstVendingShippingCallbackAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaLstVendingShippingCallbackAPIResponseModel).Reset()
}

// AlibabaLstVendingShippingCallbackAPIResponseModel is 售货机出货回传接口 成功返回结果
type AlibabaLstVendingShippingCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_lst_vending_shipping_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 执行结果
	Module bool `json:"module,omitempty" xml:"module,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaLstVendingShippingCallbackAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Module = false
}

var poolAlibabaLstVendingShippingCallbackAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaLstVendingShippingCallbackAPIResponse)
	},
}

// GetAlibabaLstVendingShippingCallbackAPIResponse 从 sync.Pool 获取 AlibabaLstVendingShippingCallbackAPIResponse
func GetAlibabaLstVendingShippingCallbackAPIResponse() *AlibabaLstVendingShippingCallbackAPIResponse {
	return poolAlibabaLstVendingShippingCallbackAPIResponse.Get().(*AlibabaLstVendingShippingCallbackAPIResponse)
}

// ReleaseAlibabaLstVendingShippingCallbackAPIResponse 将 AlibabaLstVendingShippingCallbackAPIResponse 保存到 sync.Pool
func ReleaseAlibabaLstVendingShippingCallbackAPIResponse(v *AlibabaLstVendingShippingCallbackAPIResponse) {
	v.Reset()
	poolAlibabaLstVendingShippingCallbackAPIResponse.Put(v)
}
