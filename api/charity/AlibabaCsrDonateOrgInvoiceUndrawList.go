package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCsrDonateOrgInvoiceUndrawList 获取机构待开票列表
// alibaba.csr.donate.org.invoice.undraw.list
//
// 获取机构待开票列表
func AlibabaCsrDonateOrgInvoiceUndrawList(clt *core.SDKClient, req *charity.AlibabaCsrDonateOrgInvoiceUndrawListAPIRequest, resp *charity.AlibabaCsrDonateOrgInvoiceUndrawListAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
