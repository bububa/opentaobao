package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceunitordercheck 服务商订购单上传核对
// alibaba.einvoice.unitorder.check
//
// 开票服务商回传收到的订购单用于电子发票平台核对
func Alibabaeinvoiceunitordercheck(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceunitordercheckAPIRequest, session string) (*einvoice.AlibabaeinvoiceunitordercheckAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceunitordercheckAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
