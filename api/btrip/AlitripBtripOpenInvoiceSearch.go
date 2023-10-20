package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenInvoiceSearch 差旅申请用户搜索可用发票列表
// alitrip.btrip.open.invoice.search
//
// 差旅申请用户搜索可用发票列表
func AlitripBtripOpenInvoiceSearch(clt *core.SDKClient, req *btrip.AlitripBtripOpenInvoiceSearchAPIRequest, resp *btrip.AlitripBtripOpenInvoiceSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
