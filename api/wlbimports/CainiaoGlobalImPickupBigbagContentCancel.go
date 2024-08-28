package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagContentCancel 进口大包取消
// cainiao.global.im.pickup.bigbag.content.cancel
//
// 进口大包取消
func CainiaoGlobalImPickupBigbagContentCancel(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagContentCancelAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
