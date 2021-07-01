package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripInvoiceSearch
根据发票抬头搜索发票
alitrip.btrip.invoice.search

用户根据发票抬头搜索发票信息 */
func AlitripBtripInvoiceSearch(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceSearchAPIRequest, session string) (*btrip.AlitripBtripInvoiceSearchAPIResponse, error) {
	var resp btrip.AlitripBtripInvoiceSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
