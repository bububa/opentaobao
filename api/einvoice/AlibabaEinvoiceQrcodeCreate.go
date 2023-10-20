package einvoice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// Alibabaeinvoiceqrcodecreate 扫码开票二维码生成
// alibaba.einvoice.qrcode.create
//
// 扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
func Alibabaeinvoiceqrcodecreate(clt *core.SDKClient, req *einvoice.AlibabaeinvoiceqrcodecreateAPIRequest, session string) (*einvoice.AlibabaeinvoiceqrcodecreateAPIResponse, error) {
	var resp einvoice.AlibabaeinvoiceqrcodecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
