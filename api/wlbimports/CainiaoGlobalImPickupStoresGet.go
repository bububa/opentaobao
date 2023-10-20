package wlbimports

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wlbimports"
)

// CainiaoGlobalImPickupStoresGet 首公里揽收-集货仓列表查询
// cainiao.global.im.pickup.stores.get
//
// 首公里揽收-集货仓列表查询
func CainiaoGlobalImPickupStoresGet(clt *core.SDKClient, req *wlbimports.CainiaoGlobalImPickupStoresGetAPIRequest, resp *wlbimports.CainiaoGlobalImPickupStoresGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
