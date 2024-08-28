package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemEdit 租赁商品编辑
// alibaba.idle.rent.item.edit
//
// 发布闲鱼租赁商品
func AlibabaIdleRentItemEdit(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentItemEditAPIRequest, resp *idle.AlibabaIdleRentItemEditAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
