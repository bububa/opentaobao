package fpm

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// Alibabafpmfileupload 结算单文件上传
// alibaba.fpm.file.upload
//
// 结算单文件上传
func Alibabafpmfileupload(clt *core.SDKClient, req *fpm.AlibabafpmfileuploadAPIRequest, session string) (*fpm.AlibabafpmfileuploadAPIResponse, error) {
	var resp fpm.AlibabafpmfileuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
