package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryQuery 查询商品库存信息
// taobao.inventory.query
//
// 建议使用新接口：tmall.inventory.query.forstore ，新ISV不推荐使用。
// 商家查询商品总体库存信息
func TaobaoInventoryQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryQueryAPIRequest, resp *fenxiao.TaobaoInventoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
