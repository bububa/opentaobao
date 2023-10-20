package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicebilleinvoicelist 扫码开票列表
// alibaba.einvoice.bill.einvoice.list
//
// 扫码开票列表，包括用户扫二维码开票和结算单同步前的开票数据
func Alibabaeinvoicebilleinvoicelist(clt *core.SDKClient, req *einvoice.AlibabaeinvoicebilleinvoicelistAPIRequest, session string) (*einvoice.AlibabaeinvoicebilleinvoicelistAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicebilleinvoicelistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
