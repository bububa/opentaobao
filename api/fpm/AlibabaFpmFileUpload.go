package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaFpmFileUpload 结算单文件上传
// alibaba.fpm.file.upload
//
// 结算单文件上传
func AlibabaFpmFileUpload(clt *core.SDKClient, req *fpm.AlibabaFpmFileUploadAPIRequest, session string) (*fpm.AlibabaFpmFileUploadAPIResponse, error) {
	var resp fpm.AlibabaFpmFileUploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
