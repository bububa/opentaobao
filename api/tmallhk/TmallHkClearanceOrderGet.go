package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TmallHkClearanceOrderGet 天猫国际订单清关信息
// tmall.hk.clearance.order.get
//
// 天猫国际订单清关信息
func TmallHkClearanceOrderGet(ctx context.Context, clt *core.SDKClient, req *tmallhk.TmallHkClearanceOrderGetAPIRequest, resp *tmallhk.TmallHkClearanceOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
