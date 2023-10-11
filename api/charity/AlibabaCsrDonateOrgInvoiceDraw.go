package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateOrgInvoiceDraw 机构开具商家票据信息
// alibaba.csr.donate.org.invoice.draw
//
// 机构开具商家票据信息
func AlibabaCsrDonateOrgInvoiceDraw(clt *core.SDKClient, req *charity.AlibabaCsrDonateOrgInvoiceDrawAPIRequest, session string) (*charity.AlibabaCsrDonateOrgInvoiceDrawAPIResponse, error) {
	var resp charity.AlibabaCsrDonateOrgInvoiceDrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
