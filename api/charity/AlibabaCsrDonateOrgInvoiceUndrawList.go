package charity

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// Alibabacsrdonateorginvoiceundrawlist 获取机构待开票列表
// alibaba.csr.donate.org.invoice.undraw.list
//
// 获取机构待开票列表
func Alibabacsrdonateorginvoiceundrawlist(clt *core.SDKClient, req *charity.AlibabacsrdonateorginvoiceundrawlistAPIRequest, session string) (*charity.AlibabacsrdonateorginvoiceundrawlistAPIResponse, error) {
	var resp charity.AlibabacsrdonateorginvoiceundrawlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
