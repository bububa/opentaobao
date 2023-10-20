package txcs

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/txcs"
)

// Tmalltxcsfinanceinvoiceinput 供应商发票录入
// tmall.txcs.finance.invoice.input
//
// 提供天猫超市外部合作商家财务：供应商发票录入
func Tmalltxcsfinanceinvoiceinput(clt *core.SDKClient, req *txcs.TmalltxcsfinanceinvoiceinputAPIRequest, session string) (*txcs.TmalltxcsfinanceinvoiceinputAPIResponse, error) {
	var resp txcs.TmalltxcsfinanceinvoiceinputAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
