package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryrequestAcceptpayment 受理发薪
// alibaba.einvoice.tax.opt.salaryrequest.acceptpayment
//
// 发薪受理接口
func AlibabaEinvoiceTaxOptSalaryrequestAcceptpayment(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
