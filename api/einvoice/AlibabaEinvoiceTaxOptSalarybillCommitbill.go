package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalarybillCommitbill 提交发薪账单
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
func AlibabaEinvoiceTaxOptSalarybillCommitbill(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
