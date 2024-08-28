package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateOrgInvoiceDraw 机构开具商家票据信息
// alibaba.csr.donate.org.invoice.draw
//
// 机构开具商家票据信息
func AlibabaCsrDonateOrgInvoiceDraw(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCsrDonateOrgInvoiceDrawAPIRequest, resp *charity.AlibabaCsrDonateOrgInvoiceDrawAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
