package wlbimports

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupStoresGet 首公里揽收-集货仓列表查询
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
func CainiaoGlobalImPickupStoresGet(ctx context.Context, clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupStoresGetAPIRequest, resp *wlbimports.CainiaoGlobalImPickupStoresGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
