package wlb

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlb"
)

// TaobaoWlbItemQuery 分页查询商品
// taobao.wlb.item.query
//
// 根据状态、卖家、SKU等信息查询商品列表
func TaobaoWlbItemQuery(ctx context.Context, clt *core.SDKClient, req *wlb.TaobaoWlbItemQueryAPIRequest, resp *wlb.TaobaoWlbItemQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
