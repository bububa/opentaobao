package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemMapGet 根据物流宝商品ID查询商品映射关系
// taobao.wlb.item.map.get
//
// 根据物流宝商品ID查询商品映射关系
func TaobaoWlbItemMapGet(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbItemMapGetAPIRequest, resp *wlb.TaobaoWlbItemMapGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
