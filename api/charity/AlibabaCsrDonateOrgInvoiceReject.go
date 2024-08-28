package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateOrgInvoiceReject 机构驳回商家票据信息
// alibaba.csr.donate.org.invoice.reject
//
// 机构驳回商家票据信息
func AlibabaCsrDonateOrgInvoiceReject(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCsrDonateOrgInvoiceRejectAPIRequest, resp *charity.AlibabaCsrDonateOrgInvoiceRejectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
