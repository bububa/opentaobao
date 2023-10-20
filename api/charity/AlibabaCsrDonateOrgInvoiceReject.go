package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrdonateorginvoicereject 机构驳回商家票据信息
// alibaba.csr.donate.org.invoice.reject
//
// 机构驳回商家票据信息
func Alibabacsrdonateorginvoicereject(clt *core.SDKClient, req *charity.AlibabacsrdonateorginvoicerejectAPIRequest, session string) (*charity.AlibabacsrdonateorginvoicerejectAPIResponse, error) {
	var resp charity.AlibabacsrdonateorginvoicerejectAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
