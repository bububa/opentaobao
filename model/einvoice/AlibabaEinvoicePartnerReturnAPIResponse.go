package einvoice

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoicePartnerReturnAPIResponse 开票商回传开票结果 API返回值
// alibaba.einvoice.partner.return
//
// 开票商返回开票结果数据
type AlibabaEinvoicePartnerReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoicePartnerReturnAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaEinvoicePartnerReturnAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaEinvoicePartnerReturnAPIResponseModel).Reset()
}

// AlibabaEinvoicePartnerReturnAPIResponseModel is 开票商回传开票结果 成功返回结果
type AlibabaEinvoicePartnerReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_partner_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 服务端接收开票回传数据的结果
	Result bool `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaEinvoicePartnerReturnAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = false
}

var poolAlibabaEinvoicePartnerReturnAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaEinvoicePartnerReturnAPIResponse)
	},
}

// GetAlibabaEinvoicePartnerReturnAPIResponse 从 sync.Pool 获取 AlibabaEinvoicePartnerReturnAPIResponse
func GetAlibabaEinvoicePartnerReturnAPIResponse() *AlibabaEinvoicePartnerReturnAPIResponse {
	return poolAlibabaEinvoicePartnerReturnAPIResponse.Get().(*AlibabaEinvoicePartnerReturnAPIResponse)
}

// ReleaseAlibabaEinvoicePartnerReturnAPIResponse 将 AlibabaEinvoicePartnerReturnAPIResponse 保存到 sync.Pool
func ReleaseAlibabaEinvoicePartnerReturnAPIResponse(v *AlibabaEinvoicePartnerReturnAPIResponse) {
	v.Reset()
	poolAlibabaEinvoicePartnerReturnAPIResponse.Put(v)
}
