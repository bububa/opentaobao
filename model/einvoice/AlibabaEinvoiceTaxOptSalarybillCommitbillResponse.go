package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发薪账单 API返回值 
alibaba.einvoice.tax.opt.salarybill.commitbill

提交发薪账单
*/
type AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalarybillCommitbillResponse
}

// 提交发薪账单 成功返回结果
type AlibabaEinvoiceTaxOptSalarybillCommitbillResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salarybill_commitbill_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
}
