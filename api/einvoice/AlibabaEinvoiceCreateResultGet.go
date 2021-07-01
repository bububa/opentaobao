package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

/* AlibabaEinvoiceCreateResultGet
ERP开票结果获取
alibaba.einvoice.create.result.get

ERP开票结果获取 */
func AlibabaEinvoiceCreateResultGet(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceCreateResultGetAPIRequest, session string) (*einvoice.AlibabaEinvoiceCreateResultGetAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceCreateResultGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
