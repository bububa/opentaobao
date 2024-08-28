package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecycleOrderFulfillSync 同步回收单最终履约方式
// taobao.recycle.order.fulfill.sync
//
// 同步回收单最终履约方式
func TaobaoRecycleOrderFulfillSync(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoRecycleOrderFulfillSyncAPIRequest, resp *servicecenter.TaobaoRecycleOrderFulfillSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
