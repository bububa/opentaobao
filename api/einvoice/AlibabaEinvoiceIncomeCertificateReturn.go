package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceincomecertificatereturn 服务商回传进项认证结果
// alibaba.einvoice.income.certificate.return
//
// 服务商回传客户端agent所处环境的设备列表，比如扫描仪
func Alibabaeinvoiceincomecertificatereturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceincomecertificatereturnAPIRequest, session string) (*einvoice.AlibabaeinvoiceincomecertificatereturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceincomecertificatereturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
