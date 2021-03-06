package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceMerchantCreatereq 商家自研ERP开票请求接口
// alibaba.einvoice.merchant.createreq
//
// 商家自研ERP发起开票请求，无需授权，API只能使用商家入驻的税号进行开票
func AlibabaEinvoiceMerchantCreatereq(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceMerchantCreatereqAPIRequest, session string) (*einvoice.AlibabaEinvoiceMerchantCreatereqAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceMerchantCreatereqAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
