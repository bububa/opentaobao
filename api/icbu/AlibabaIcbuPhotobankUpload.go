package icbu

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbu"
)

// AlibabaIcbuPhotobankUpload 图片银行图片上传开放接口
// alibaba.icbu.photobank.upload
//
// 图片银行图片上传开放接口
func AlibabaIcbuPhotobankUpload(ctx context.Context, clt *core.SDKClient, req *icbu.AlibabaIcbuPhotobankUploadAPIRequest, resp *icbu.AlibabaIcbuPhotobankUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
