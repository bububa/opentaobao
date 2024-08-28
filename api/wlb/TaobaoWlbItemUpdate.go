package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemUpdate 物流宝商品修改
// taobao.wlb.item.update
//
// 修改物流宝商品信息
func TaobaoWlbItemUpdate(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbItemUpdateAPIRequest, resp *wlb.TaobaoWlbItemUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
