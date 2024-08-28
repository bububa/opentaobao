package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemCombinationGet 根据商品id查询商品组合关系
// taobao.wlb.item.combination.get
//
// 根据商品id查询商品组合关系
func TaobaoWlbItemCombinationGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbItemCombinationGetAPIRequest, resp *wlb.TaobaoWlbItemCombinationGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
