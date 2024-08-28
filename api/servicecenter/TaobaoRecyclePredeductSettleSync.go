package servicecenter

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/servicecenter"
)

// TaobaoRecyclePredeductSettleSync 同步回收单线下打款明细
// taobao.recycle.prededuct.settle.sync
//
// 同步回收单线下打款明细
func TaobaoRecyclePredeductSettleSync(ctx context.Context, clt *core.SDKClient, req *servicecenter.TaobaoRecyclePredeductSettleSyncAPIRequest, resp *servicecenter.TaobaoRecyclePredeductSettleSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
