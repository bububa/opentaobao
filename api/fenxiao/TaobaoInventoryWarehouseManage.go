package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryWarehouseManage 创建商家仓或者更新商家仓信息
// taobao.inventory.warehouse.manage
//
// 创建商家仓或者更新商家仓信息
func TaobaoInventoryWarehouseManage(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryWarehouseManageAPIRequest, resp *fenxiao.TaobaoInventoryWarehouseManageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
