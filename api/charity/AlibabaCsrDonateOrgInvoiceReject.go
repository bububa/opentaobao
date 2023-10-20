package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateOrgInvoiceReject 机构驳回商家票据信息
// alibaba.csr.donate.org.invoice.reject
//
// 机构驳回商家票据信息
func AlibabaCsrDonateOrgInvoiceReject(clt *core.SDKClient, req *charity.AlibabaCsrDonateOrgInvoiceRejectAPIRequest, resp *charity.AlibabaCsrDonateOrgInvoiceRejectAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
