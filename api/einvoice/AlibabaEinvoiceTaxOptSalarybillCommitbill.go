package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalarybillCommitbill 提交发薪账单
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
func AlibabaEinvoiceTaxOptSalarybillCommitbill(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
