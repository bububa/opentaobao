package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// TaobaoWlbImportsOrderCancel 一般进口取消物流订单
// taobao.wlb.imports.order.cancel
//
// 商家在发货后，需要对订单进行取消，如果仓库已经收货则无法取消。
func TaobaoWlbImportsOrderCancel(ctx context.Context, clt *core.SDKClient, req *wlbimports.TaobaoWlbImportsOrderCancelAPIRequest, resp *wlbimports.TaobaoWlbImportsOrderCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
