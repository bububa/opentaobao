package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupBigbagInfo 大包状态查询
// cainiao.global.im.pickup.bigbag.info
//
// 大包状态查询
func CainiaoGlobalImPickupBigbagInfo(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupBigbagInfoAPIRequest, resp *wlbimports.CainiaoGlobalImPickupBigbagInfoAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
