package fpm

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fpm"
)

// AlibabaFpmFileUpload 结算单文件上传
// alibaba.fpm.file.upload
//
// 结算单文件上传
func AlibabaFpmFileUpload(ctx context.Context, clt *core.SDKClient, req *fpm.AlibabaFpmFileUploadAPIRequest, resp *fpm.AlibabaFpmFileUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
