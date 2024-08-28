package idle

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentItemQuery 查询租赁商品信息
// alibaba.idle.rent.item.query
//
// 查询租赁商品信息
func AlibabaIdleRentItemQuery(ctx context.Context, clt *core.SDKClient, req *idle.AlibabaIdleRentItemQueryAPIRequest, resp *idle.AlibabaIdleRentItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
