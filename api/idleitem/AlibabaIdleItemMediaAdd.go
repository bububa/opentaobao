package idleitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleitem"
)

// AlibabaIdleItemMediaAdd 图片上传
// alibaba.idle.item.media.add
//
// 上传图片
func AlibabaIdleItemMediaAdd(ctx context.Context, clt *core.SDKClient, req *idleitem.AlibabaIdleItemMediaAddAPIRequest, resp *idleitem.AlibabaIdleItemMediaAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
