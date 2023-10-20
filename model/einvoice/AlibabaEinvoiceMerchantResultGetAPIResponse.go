package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceMerchantResultGetAPIResponse 商家自研ERP开票结果获取 API返回值
// alibaba.einvoice.merchant.result.get
//
// 商家自研ERP开票结果获取
type AlibabaEinvoiceMerchantResultGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceMerchantResultGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceMerchantResultGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceMerchantResultGetAPIResponseModel).Reset()
}

// AlibabaEinvoiceMerchantResultGetAPIResponseModel is 商家自研ERP开票结果获取 成功返回结果
type AlibabaEinvoiceMerchantResultGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_merchant_result_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票返回结果数据列表
	InvoiceResultList []InvoiceResult `json:"invoice_result_list,omitempty" xml:"invoice_result_list>invoice_result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceMerchantResultGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.InvoiceResultList = m.InvoiceResultList[:0]
}

var poolAlibabaEinvoiceMerchantResultGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceMerchantResultGetAPIResponse)
	},
}

// GetAlibabaEinvoiceMerchantResultGetAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceMerchantResultGetAPIResponse
func GetAlibabaEinvoiceMerchantResultGetAPIResponse() *AlibabaEinvoiceMerchantResultGetAPIResponse {
	return poolAlibabaEinvoiceMerchantResultGetAPIResponse.Get().(*AlibabaEinvoiceMerchantResultGetAPIResponse)
}

// ReleaseAlibabaEinvoiceMerchantResultGetAPIResponse 将 AlibabaEinvoiceMerchantResultGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceMerchantResultGetAPIResponse(v *AlibabaEinvoiceMerchantResultGetAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceMerchantResultGetAPIResponse.Put(v)
}
