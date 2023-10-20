package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryresultQuery 查询发薪结果
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
func AlibabaEinvoiceTaxOptSalaryresultQuery(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest, session string) (*einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
