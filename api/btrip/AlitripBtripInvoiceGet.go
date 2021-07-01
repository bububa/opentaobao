package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

/* AlitripBtripInvoiceGet
获取用户可用发票列表
alitrip.btrip.invoice.get

差旅申请用户获取可用发票列表 */
func AlitripBtripInvoiceGet(clt *core.SDKClient, req *btrip.AlitripBtripInvoiceGetAPIRequest, session string) (*btrip.AlitripBtripInvoiceGetAPIResponse, error) {
	var resp btrip.AlitripBtripInvoiceGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
