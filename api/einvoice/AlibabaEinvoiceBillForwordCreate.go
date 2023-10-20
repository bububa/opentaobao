package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicebillforwordcreate 扫码开票结算单同步前开发票
// alibaba.einvoice.bill.forword.create
//
// 扫码开票结算单同步前开发票，会将数据同步到结算单中
func Alibabaeinvoicebillforwordcreate(clt *core.SDKClient, req *einvoice.AlibabaeinvoicebillforwordcreateAPIRequest, session string) (*einvoice.AlibabaeinvoicebillforwordcreateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicebillforwordcreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
