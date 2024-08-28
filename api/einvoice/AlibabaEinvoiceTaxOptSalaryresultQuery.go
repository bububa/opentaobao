package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryresultQuery 查询发薪结果
// alibaba.einvoice.tax.opt.salaryresult.query
//
// 查询发薪结果
func AlibabaEinvoiceTaxOptSalaryresultQuery(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalaryresultQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
