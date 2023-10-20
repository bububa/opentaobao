package einvoice

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse 提交发薪账单 API返回值
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
type AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse struct {
	model.CommonResponse
	AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponseModel
}

// AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponseModel is 提交发薪账单 成功返回结果
type AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salarybill_commitbill_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 承包商编码
	ContractorCode string `json:"contractor_code,omitempty" xml:"contractor_code,omitempty"`
	// 明细id
	DetailId string `json:"detail_id,omitempty" xml:"detail_id,omitempty"`
	// 业务方编码
	EmployerCode string `json:"employer_code,omitempty" xml:"employer_code,omitempty"`
	// 用户在业务方平台userid
	IdentificationInBelongingEmployer string `json:"identification_in_belonging_employer,omitempty" xml:"identification_in_belonging_employer,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误原因
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}
