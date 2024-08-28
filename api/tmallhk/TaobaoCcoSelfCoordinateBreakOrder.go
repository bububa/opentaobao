package tmallhk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallhk"
)

// TaobaoCcoSelfCoordinateBreakOrder 天猫国际直购供应商毁单通知
// taobao.cco.self.coordinate.break.order
//
// 天猫国际直购供应商毁单通知
func TaobaoCcoSelfCoordinateBreakOrder(ctx context.Context, clt *core.SDKClient, req *tmallhk.TaobaoCcoSelfCoordinateBreakOrderAPIRequest, resp *tmallhk.TaobaoCcoSelfCoordinateBreakOrderAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
