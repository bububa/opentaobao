package einvoice

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
提交发薪账单 APIResponse
alibaba.einvoice.tax.opt.salarybill.commitbill

提交发薪账单
*/
type AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse struct {
    model.CommonResponse
    AlibabaEinvoiceTaxOptSalarybillCommitbillResponse
}

type AlibabaEinvoiceTaxOptSalarybillCommitbillResponse struct {
    XMLName xml.Name `xml:"alibaba_einvoice_tax_opt_salarybill_commitbill_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

}
