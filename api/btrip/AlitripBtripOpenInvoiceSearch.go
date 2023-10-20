package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripOpenInvoiceSearch 差旅申请用户搜索可用发票列表
// alitrip.btrip.open.invoice.search
//
// 差旅申请用户搜索可用发票列表
func AlitripBtripOpenInvoiceSearch(clt *core.SDKClient, req *btrip.AlitripBtripOpenInvoiceSearchAPIRequest, session string) (*btrip.AlitripBtripOpenInvoiceSearchAPIResponse, error) {
	var resp btrip.AlitripBtripOpenInvoiceSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
