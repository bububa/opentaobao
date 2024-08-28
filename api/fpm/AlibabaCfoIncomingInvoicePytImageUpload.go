package fpm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaCfoIncomingInvoicePytImageUpload 票易通发票影像上传
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
func AlibabaCfoIncomingInvoicePytImageUpload(ctx context.Context, clt *core.SDKClient, req *fpm.AlibabaCfoIncomingInvoicePytImageUploadAPIRequest, resp *fpm.AlibabaCfoIncomingInvoicePytImageUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
