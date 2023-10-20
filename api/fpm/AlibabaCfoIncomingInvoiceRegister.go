package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaCfoIncomingInvoiceRegister 发票登记接口
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
func AlibabaCfoIncomingInvoiceRegister(clt *core.SDKClient, req *fpm.AlibabaCfoIncomingInvoiceRegisterAPIRequest, session string) (*fpm.AlibabaCfoIncomingInvoiceRegisterAPIResponse, error) {
	var resp fpm.AlibabaCfoIncomingInvoiceRegisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
