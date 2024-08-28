package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceTaxOptSalarybillCommitbill 提交发薪账单
// alibaba.einvoice.tax.opt.salarybill.commitbill
//
// 提交发薪账单
func AlibabaEinvoiceTaxOptSalarybillCommitbill(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIRequest, resp *einvoice.AlibabaEinvoiceTaxOptSalarybillCommitbillAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
