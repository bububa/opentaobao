package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicemerchantresultget 商家自研ERP开票结果获取
// alibaba.einvoice.merchant.result.get
//
// 商家自研ERP开票结果获取
func Alibabaeinvoicemerchantresultget(clt *core.SDKClient, req *einvoice.AlibabaeinvoicemerchantresultgetAPIRequest, session string) (*einvoice.AlibabaeinvoicemerchantresultgetAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicemerchantresultgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
