package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalaryaccountUpdate 更新用户发薪资产
// alibaba.einvoice.tax.opt.salaryaccount.update
//
// 更新用户的发薪账号
func AlibabaEinvoiceTaxOptSalaryaccountUpdate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalaryaccountUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
