package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicetaxoptsalaryrequestacceptpayment 受理发薪
// alibaba.einvoice.tax.opt.salaryrequest.acceptpayment
//
// 发薪受理接口
func Alibabaeinvoicetaxoptsalaryrequestacceptpayment(clt *core.SDKClient, req *einvoice.AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIRequest, session string) (*einvoice.AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicetaxoptsalaryrequestacceptpaymentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
