package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripInvoiceSearch 根据发票抬头搜索发票
// alitrip.btrip.invoice.search
//
// 用户根据发票抬头搜索发票信息
func AlitripBtripInvoiceSearch(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSearchAPIRequest, resp *btrip.AlitripBtripInvoiceSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
