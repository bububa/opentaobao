package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceDeductGetAPIResponse
发票扣减的接口 API返回值
alibaba.einvoice.deduct.get

获取历史发票扣减量、每日发票扣减量的接口 */
type AlibabaEinvoiceDeductGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceDeductGetAPIResponseModel
}

// AlibabaEinvoiceDeductGetAPIResponseModel is 发票扣减的接口 成功返回结果
type AlibabaEinvoiceDeductGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_deduct_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *AlibabaEinvoiceDeductGetResultSet `json:"result,omitempty" xml:"result,omitempty"`
}
