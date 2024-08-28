package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbInventoryDetailGet 查询库存详情
// taobao.wlb.inventory.detail.get
//
// 查询库存详情，通过商品ID获取发送请求的卖家的库存详情
func TaobaoWlbInventoryDetailGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbInventoryDetailGetAPIRequest, resp *wlb.TaobaoWlbInventoryDetailGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
