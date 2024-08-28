package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvMediaUpload 闲鱼服务商-图片上传
// alibaba.idle.isv.media.upload
//
// 供外部服务商ISV进行闲鱼商品发布时上传商品所需图片
func AlibabaIdleIsvMediaUpload(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvMediaUploadAPIRequest, resp *idleisv.AlibabaIdleIsvMediaUploadAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
