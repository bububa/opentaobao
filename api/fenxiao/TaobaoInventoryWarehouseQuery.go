package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryWarehouseQuery 分页查询商家仓信息
// taobao.inventory.warehouse.query
//
// 分页查询商家仓信息
func TaobaoInventoryWarehouseQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryWarehouseQueryAPIRequest, resp *fenxiao.TaobaoInventoryWarehouseQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
