package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryaccountQuery 查询用户的发薪账号
// alibaba.einvoice.tax.opt.salaryaccount.query
//
// 查询用户的发薪账号状态
func AlibabaEinvoiceTaxOptSalaryaccountQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryaccountQueryAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalaryaccountQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
