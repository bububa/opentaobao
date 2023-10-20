package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceUnitorderCheck 服务商订购单上传核对
// alibaba.einvoice.unitorder.check
//
// 开票服务商回传收到的订购单用于电子发票平台核对
func AlibabaEinvoiceUnitorderCheck(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceUnitorderCheckAPIRequest, resp *einvoice.AlibabaEinvoiceUnitorderCheckAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
