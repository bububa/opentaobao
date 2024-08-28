package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryInitialItem 商品库存初始化
// taobao.inventory.initial.item
//
// 建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
// 商家仓商品初始化在各个仓中库存
func TaobaoInventoryInitialItem(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryInitialItemAPIRequest, resp *fenxiao.TaobaoInventoryInitialItemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
