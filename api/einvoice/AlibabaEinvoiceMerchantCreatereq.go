package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicemerchantcreatereq 商家自研ERP开票请求接口
// alibaba.einvoice.merchant.createreq
//
// 商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
func Alibabaeinvoicemerchantcreatereq(clt *core.SDKClient, req *einvoice.AlibabaeinvoicemerchantcreatereqAPIRequest, session string) (*einvoice.AlibabaeinvoicemerchantcreatereqAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicemerchantcreatereqAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
