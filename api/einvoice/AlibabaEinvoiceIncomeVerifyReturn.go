package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceincomeverifyreturn 服务商回传发票查验的结果
// alibaba.einvoice.income.verify.return
//
// 服务商回传发票查验的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的查验回传
func Alibabaeinvoiceincomeverifyreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceincomeverifyreturnAPIRequest, session string) (*einvoice.AlibabaeinvoiceincomeverifyreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceincomeverifyreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
