package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsInventoryProfitlossGet 通过订单列表批量获取库存损益单信息
// taobao.wlb.wms.inventory.profitloss.get
//
// 通过订单列表批量获取库存损益单信息
func TaobaoWlbWmsInventoryProfitlossGet(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsInventoryProfitlossGetAPIRequest, resp *wms.TaobaoWlbWmsInventoryProfitlossGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
