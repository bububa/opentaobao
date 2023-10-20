package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceincomeocrreturn 服务商回传发票ocr的结果
// alibaba.einvoice.income.ocr.return
//
// 服务商回传发票ocr的结果，分两种场景：扫描驱动服务商主动回传；阿里主动发起的ocr回传
func Alibabaeinvoiceincomeocrreturn(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceincomeocrreturnAPIRequest, session string) (*einvoice.AlibabaeinvoiceincomeocrreturnAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceincomeocrreturnAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
