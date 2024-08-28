package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsOrderGet 物流订单获取
// taobao.wlb.imports.order.get
//
// 一般进口物流订单获取
func TaobaoWlbImportsOrderGet(ctx context.Context, clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsOrderGetAPIRequest, resp *wlbimports.TaobaoWlbImportsOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
