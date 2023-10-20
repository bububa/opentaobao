package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrdonateorginvoicedraw 机构开具商家票据信息
// alibaba.csr.donate.org.invoice.draw
//
// 机构开具商家票据信息
func Alibabacsrdonateorginvoicedraw(clt *core.SDKClient, req *charity.AlibabacsrdonateorginvoicedrawAPIRequest, session string) (*charity.AlibabacsrdonateorginvoicedrawAPIResponse, error) {
	var resp charity.AlibabacsrdonateorginvoicedrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
