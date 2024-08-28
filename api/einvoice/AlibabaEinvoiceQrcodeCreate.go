package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoiceQrcodeCreate 扫码开票二维码生成
// alibaba.einvoice.qrcode.create
//
// 扫码开票功能中的二维码生成接口，pos机等发起请求生成二维码
func AlibabaEinvoiceQrcodeCreate(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoiceQrcodeCreateAPIRequest, resp *einvoice.AlibabaEinvoiceQrcodeCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
