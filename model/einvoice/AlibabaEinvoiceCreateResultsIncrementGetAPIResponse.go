package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceCreateResultsIncrementGetAPIResponse ERP增量开票结果获取 API返回值
// alibaba.einvoice.create.results.increment.get
//
// 增量开票结果获取
type AlibabaEinvoiceCreateResultsIncrementGetAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceCreateResultsIncrementGetAPIResponseModel
}

// AlibabaEinvoiceCreateResultsIncrementGetAPIResponseModel is ERP增量开票结果获取 成功返回结果
type AlibabaEinvoiceCreateResultsIncrementGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_create_results_increment_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 开票结果返回列表
	InvoiceResultList []InvoiceResult `json:"invoice_result_list,omitempty" xml:"invoice_result_list>invoice_result,omitempty"`
	// 符合条件的开票总数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}
