package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateOrgInvoiceDraw 机构开具商家票据信息
// alibaba.csr.donate.org.invoice.draw
//
// 机构开具商家票据信息
func AlibabaCsrDonateOrgInvoiceDraw(clt *core.SDKClient, req *charity.AlibabaCsrDonateOrgInvoiceDrawAPIRequest, resp *charity.AlibabaCsrDonateOrgInvoiceDrawAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
