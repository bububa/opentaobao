package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryrequestAcceptpayment 受理发薪
// alibaba.einvoice.tax.opt.salaryrequest.acceptpayment
//
// 发薪受理接口
func AlibabaEinvoiceTaxOptSalaryrequestAcceptpayment(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceTaxOptSalaryrequestAcceptpaymentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
