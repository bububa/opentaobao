package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaCfoIncomingInvoicePytImageUpload 票易通发票影像上传
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
func AlibabaCfoIncomingInvoicePytImageUpload(clt *core.SDKClient, req *fpm.AlibabaCfoIncomingInvoicePytImageUploadAPIRequest, session string) (*fpm.AlibabaCfoIncomingInvoicePytImageUploadAPIResponse, error) {
	var resp fpm.AlibabaCfoIncomingInvoicePytImageUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
