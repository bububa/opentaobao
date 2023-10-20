package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryresultQuery 查询发薪结果
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
func AlibabaEinvoiceTaxOptSalaryresultQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
