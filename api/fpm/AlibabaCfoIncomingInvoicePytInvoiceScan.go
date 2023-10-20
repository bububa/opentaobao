package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaCfoIncomingInvoicePytInvoiceScan 票易通发票ocr信息同步
// alibaba.cfo.incoming.invoice.pyt.invoice.scan
//
// 票易通发票ocr信息同步
func AlibabaCfoIncomingInvoicePytInvoiceScan(clt *core.SDKClient, req *fpm.AlibabaCfoIncomingInvoicePytInvoiceScanAPIRequest, resp *fpm.AlibabaCfoIncomingInvoicePytInvoiceScanAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
