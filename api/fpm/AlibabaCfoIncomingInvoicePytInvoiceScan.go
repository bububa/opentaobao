package fpm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaCfoIncomingInvoicePytInvoiceScan 票易通发票ocr信息同步
// alibaba.cfo.incoming.invoice.pyt.invoice.scan
//
// 票易通发票ocr信息同步
func AlibabaCfoIncomingInvoicePytInvoiceScan(ctx context.Context, clt *core.SDKClient, req *fpm.AlibabaCfoIncomingInvoicePytInvoiceScanAPIRequest, resp *fpm.AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
