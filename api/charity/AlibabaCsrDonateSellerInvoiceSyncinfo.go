package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrdonatesellerinvoicesyncinfo 链上同步商家票据信息
// alibaba.csr.donate.seller.invoice.syncinfo
//
// 链上同步商家票据信息
func Alibabacsrdonatesellerinvoicesyncinfo(clt *core.SDKClient, req *charity.AlibabacsrdonatesellerinvoicesyncinfoAPIRequest, session string) (*charity.AlibabacsrdonatesellerinvoicesyncinfoAPIResponse, error) {
	var resp charity.AlibabacsrdonatesellerinvoicesyncinfoAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
