package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemAdd 租赁商品发布
// alibaba.idle.rent.item.add
//
// 发布闲鱼租赁商品
func AlibabaIdleRentItemAdd(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentItemAddAPIRequest, resp *idle.AlibabaIdleRentItemAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
