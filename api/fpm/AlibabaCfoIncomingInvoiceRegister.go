package fpm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaCfoIncomingInvoiceRegister 发票登记接口
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
func AlibabaCfoIncomingInvoiceRegister(ctx context.Context, clt *core.SDKClient, req *fpm.AlibabaCfoIncomingInvoiceRegisterAPIRequest, resp *fpm.AlibabaCfoIncomingInvoiceRegisterAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
