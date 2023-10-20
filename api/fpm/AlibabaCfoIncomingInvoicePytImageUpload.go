package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// Alibabacfoincominginvoicepytimageupload 票易通发票影像上传
// alibaba.cfo.incoming.invoice.pyt.image.upload
//
// 票易通发票影像上传
func Alibabacfoincominginvoicepytimageupload(clt *core.SDKClient, req *fpm.AlibabacfoincominginvoicepytimageuploadAPIRequest, session string) (*fpm.AlibabacfoincominginvoicepytimageuploadAPIResponse, error) {
	var resp fpm.AlibabacfoincominginvoicepytimageuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
