package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// Alibabacfoincominginvoiceregister 发票登记接口
// alibaba.cfo.incoming.invoice.register
//
// 发票登记接口
func Alibabacfoincominginvoiceregister(clt *core.SDKClient, req *fpm.AlibabacfoincominginvoiceregisterAPIRequest, session string) (*fpm.AlibabacfoincominginvoiceregisterAPIResponse, error) {
	var resp fpm.AlibabacfoincominginvoiceregisterAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
