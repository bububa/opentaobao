package tax

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabataxinvoicesyncAPIResponse 第三方开票回调API API返回值
// alibaba.tax.invoice.sync
//
// 该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
type AlibabataxinvoicesyncAPIResponse struct {
	model.CommonResponse
	AlibabataxinvoicesyncAPIResponseModel
}

// AlibabataxinvoicesyncAPIResponseModel is 第三方开票回调API 成功返回结果
type AlibabataxinvoicesyncAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_tax_invoice_sync_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Result *ThirdPartyInvoiceCallBackResultDto `json:"result,omitempty" xml:"result,omitempty"`
}
