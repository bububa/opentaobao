package einvoice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/einvoice"
)

// AlibabaEinvoicePartnerUpload 服务商发票上传接口（非授权）
// alibaba.einvoice.partner.upload
//
// 服务商发票上传接口（非授权）
func AlibabaEinvoicePartnerUpload(ctx context.Context, clt *core.SDKClient, req *einvoice.AlibabaEinvoicePartnerUploadAPIRequest, resp *einvoice.AlibabaEinvoicePartnerUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
