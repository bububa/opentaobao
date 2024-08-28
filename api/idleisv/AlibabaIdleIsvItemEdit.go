package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvItemEdit 服务商闲鱼商品编辑
// alibaba.idle.isv.item.edit
//
// 服务商ISV闲鱼商品编辑操作
func AlibabaIdleIsvItemEdit(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvItemEditAPIRequest, resp *idleisv.AlibabaIdleIsvItemEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
