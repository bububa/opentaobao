package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceIncomeOcrReturnAPIResponse 服务商回传发票ocr的结果 API返回值
// alibaba.einvoice.income.ocr.return
//
// 服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
type AlibabaEinvoiceIncomeOcrReturnAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceIncomeOcrReturnAPIResponseModel
}

// AlibabaEinvoiceIncomeOcrReturnAPIResponseModel is 服务商回传发票ocr的结果 成功返回结果
type AlibabaEinvoiceIncomeOcrReturnAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_income_ocr_return_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口调用结果，true=成功，false=失败，subCode以isp开头时需要服务商重试
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
