package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceMerchantDeleteAPIResponse 发票中台-同平台取消授权税号适用商户 API返回值
// alibaba.einvoice.merchant.delete
//
// 税号授权给同平台下其他商户使用后，可以使用此接口取消授权，被取消授权的商户失去开票能力
type AlibabaEinvoiceMerchantDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceMerchantDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceMerchantDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceMerchantDeleteAPIResponseModel).Reset()
}

// AlibabaEinvoiceMerchantDeleteAPIResponseModel is 发票中台-同平台取消授权税号适用商户 成功返回结果
type AlibabaEinvoiceMerchantDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_merchant_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 删除成功的业务平台商户ID
	MerchantUserId string `json:"merchant_user_id,omitempty" xml:"merchant_user_id,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceMerchantDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.MerchantUserId = ""
}

var poolAlibabaEinvoiceMerchantDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceMerchantDeleteAPIResponse)
	},
}

// GetAlibabaEinvoiceMerchantDeleteAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceMerchantDeleteAPIResponse
func GetAlibabaEinvoiceMerchantDeleteAPIResponse() *AlibabaEinvoiceMerchantDeleteAPIResponse {
	return poolAlibabaEinvoiceMerchantDeleteAPIResponse.Get().(*AlibabaEinvoiceMerchantDeleteAPIResponse)
}

// ReleaseAlibabaEinvoiceMerchantDeleteAPIResponse 将 AlibabaEinvoiceMerchantDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceMerchantDeleteAPIResponse(v *AlibabaEinvoiceMerchantDeleteAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceMerchantDeleteAPIResponse.Put(v)
}
