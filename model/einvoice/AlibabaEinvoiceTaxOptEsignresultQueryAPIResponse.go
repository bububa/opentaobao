package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse
查询用户签约税优结果 API返回值
alibaba.einvoice.tax.opt.esignresult.query

查询用户是否已经签约 */
type AlibabaEinvoiceTaxOptEsignresultQueryAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel
}

// AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel is 查询用户签约税优结果 成功返回结果
type AlibabaEinvoiceTaxOptEsignresultQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_esignresult_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 查询结果
	Results []AgreementInfoDto `json:"results,omitempty" xml:"results>agreement_info_dto,omitempty"`
}
