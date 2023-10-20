package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceMerchantCreatereqAPIResponse 商家自研ERP开票请求接口 API返回值
// alibaba.einvoice.merchant.createreq
//
// 商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
type AlibabaEinvoiceMerchantCreatereqAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceMerchantCreatereqAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoiceMerchantCreatereqAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoiceMerchantCreatereqAPIResponseModel).Reset()
}

// AlibabaEinvoiceMerchantCreatereqAPIResponseModel is 商家自研ERP开票请求接口 成功返回结果
type AlibabaEinvoiceMerchantCreatereqAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_merchant_createreq_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票信息是否成功接受
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoiceMerchantCreatereqAPIResponseModel) Reset() {
	m.RequestId = ""
	m.IsSuccess = false
}

var poolAlibabaEinvoiceMerchantCreatereqAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoiceMerchantCreatereqAPIResponse)
	},
}

// GetAlibabaEinvoiceMerchantCreatereqAPIResponse 从 sync.Pool 获取 AlibabaEinvoiceMerchantCreatereqAPIResponse
func GetAlibabaEinvoiceMerchantCreatereqAPIResponse() *AlibabaEinvoiceMerchantCreatereqAPIResponse {
	return poolAlibabaEinvoiceMerchantCreatereqAPIResponse.Get().(*AlibabaEinvoiceMerchantCreatereqAPIResponse)
}

// ReleaseAlibabaEinvoiceMerchantCreatereqAPIResponse 将 AlibabaEinvoiceMerchantCreatereqAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoiceMerchantCreatereqAPIResponse(v *AlibabaEinvoiceMerchantCreatereqAPIResponse) {
	v.Reset()
	poolAlibabaEinvoiceMerchantCreatereqAPIResponse.Put(v)
}
