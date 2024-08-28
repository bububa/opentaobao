package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagContentCreate 大包创建
// cainiao.global.im.pickup.bigbag.content.create
//
// 大包创建
func CainiaoGlobalImPickupBigbagContentCreate(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagContentCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
