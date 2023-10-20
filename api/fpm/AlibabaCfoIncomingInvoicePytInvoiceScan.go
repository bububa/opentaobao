package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// Alibabacfoincominginvoicepytinvoicescan 票易通发票ocr信息同步
// alibaba.cfo.incoming.invoice.pyt.invoice.scan
//
// 票易通发票ocr信息同步
func Alibabacfoincominginvoicepytinvoicescan(clt *core.SDKClient, req *fpm.AlibabacfoincominginvoicepytinvoicescanAPIRequest, session string) (*fpm.AlibabacfoincominginvoicepytinvoicescanAPIResponse, error) {
	var resp fpm.AlibabacfoincominginvoicepytinvoicescanAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
