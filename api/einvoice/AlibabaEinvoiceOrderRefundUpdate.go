package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceorderrefundupdate 回传订单退款审核结果
// alibaba.einvoice.order.refund.update
//
// ISV回传订单退款审核结果
func Alibabaeinvoiceorderrefundupdate(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceorderrefundupdateAPIRequest, session string) (*einvoice.AlibabaeinvoiceorderrefundupdateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceorderrefundupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
