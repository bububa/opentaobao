package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceIncomeCertificateReturn 服务商回传进项认证结果
// alibaba.einvoice.income.certificate.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
func AlibabaEinvoiceIncomeCertificateReturn(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceIncomeCertificateReturnAPIRequest, resp *einvoice.AlibabaEinvoiceIncomeCertificateReturnAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
