package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvItemPublish 服务商闲鱼商品发布
// alibaba.idle.isv.item.publish
//
// 服务商ISV闲鱼商品发布
func AlibabaIdleIsvItemPublish(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemPublishAPIRequest, resp *idleisv.AlibabaIdleIsvItemPublishAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
