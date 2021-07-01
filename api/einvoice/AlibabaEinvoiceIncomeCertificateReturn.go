package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

/* AlibabaEinvoiceIncomeCertificateReturn
服务商回传进项认证结果
alibaba.einvoice.income.certificate.return

服务商回传客户端agent所处环境的设备列表，比如扫描仪 */
func AlibabaEinvoiceIncomeCertificateReturn(clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeCertificateReturnAPIRequest, session string) (*einvoice.AlibabaEinvoiceIncomeCertificateReturnAPIResponse, error) {
	var resp einvoice.AlibabaEinvoiceIncomeCertificateReturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
