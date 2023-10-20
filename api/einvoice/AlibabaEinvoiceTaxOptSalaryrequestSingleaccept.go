package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryrequestSingleaccept 单明细发薪受理
// alibaba.einvoice.tax.opt.salaryrequest.singleaccept
//
// 单明细发薪受理
func AlibabaEinvoiceTaxOptSalaryrequestSingleaccept(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalaryrequestSingleacceptAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
