package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateSellerInvoiceSyncinfo 链上同步商家票据信息
// alibaba.csr.donate.seller.invoice.syncinfo
//
// 链上同步商家票据信息
func AlibabaCsrDonateSellerInvoiceSyncinfo(clt *core.SDKClient, req *charity.AlibabaCsrDonateSellerInvoiceSyncinfoAPIRequest, session string) (*charity.AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse, error) {
	var resp charity.AlibabaCsrDonateSellerInvoiceSyncinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
