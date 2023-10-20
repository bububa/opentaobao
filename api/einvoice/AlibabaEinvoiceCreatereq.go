package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoicecreatereq ERP开票请求接口
// alibaba.einvoice.createreq
//
// ERP发起开票请求
func Alibabaeinvoicecreatereq(clt *core.SDKClient, req *einvoice.AlibabaeinvoicecreatereqAPIRequest, session string) (*einvoice.AlibabaeinvoicecreatereqAPIResponse, error) {
	var resp einvoice.AlibabaeinvoicecreatereqAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
