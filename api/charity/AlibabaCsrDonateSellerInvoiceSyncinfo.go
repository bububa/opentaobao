package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateSellerInvoiceSyncinfo 链上同步商家票据信息
// alibaba.csr.donate.seller.invoice.syncinfo
//
// 链上同步商家票据信息
func AlibabaCsrDonateSellerInvoiceSyncinfo(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest, resp *charity.AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
